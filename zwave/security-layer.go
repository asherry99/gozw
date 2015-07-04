package zwave

import (
	"encoding/hex"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/bjyoungblood/gozw/zwave/commandclass"
)

const (
	InternalNonceTTL           = time.Second * 15
	ExternalNonceTTL           = time.Second * 10
	MaxSecureInclusionDuration = time.Second * 30
	NonceRequestTimeout        = time.Second * 10
)

// Note: always use the smallest size based on the options
// @todo: also, implement the ability to use different options
const (
	SecurePayloadMaxSizeExplore   = 26
	SecurePayloadMaxSizeAutoRoute = 28
	SecurePayloadMaxSizeNoRoute   = 34
)

const (
	SecuritySequenceSequencedFlag   uint8 = 0x10
	SecuritySequenceSecondFrameFlag       = 0x20
	SecuritySequenceCounterMask           = 0x0f
)

const (
	SecuritySequenceCounterMin uint8 = 1
	SecuritySequenceCounterMax       = 15
)

type SecurityLayer struct {
	session *SessionLayer

	// internal nonce table is keyed by the first byte of the nonce
	internalNonceTable *NonceTable

	// external nonce table is keyed by the node id
	externalNonceTable *NonceTable

	// maps node id to
	waitForNonce map[byte]chan bool
	waitMapLock  *sync.Mutex

	// one sequence counter per node; used to properly pair sequenced security messages
	sequenceCounters map[uint8]uint8

	// whether we are currently performing secure inclusion of a node, or just
	// processing normally
	secureInclusionMode     bool
	secureInclusionComplete chan bool
	secureInclusionStage    chan int
	includingNode           *Node
}

func NewSecurityLayer(session *SessionLayer) *SecurityLayer {
	securityLayer := &SecurityLayer{
		session: session,

		internalNonceTable: NewNonceTable(),
		externalNonceTable: NewNonceTable(),

		waitForNonce: map[byte]chan bool{},
		waitMapLock:  &sync.Mutex{}, // @todo add locking

		sequenceCounters: map[uint8]uint8{},

		secureInclusionMode:     false,
		secureInclusionComplete: make(chan bool),
		secureInclusionStage:    make(chan int),
		includingNode:           nil,
	}

	return securityLayer
}

func (s *SecurityLayer) SecurityFrameHandler(cmd *ApplicationCommandHandlerBridge, frame *Frame) {

	switch cmd.CommandData[1] {
	case commandclass.CommandSecurityVersion:
		fmt.Println("Received SecurityVersion")

	case commandclass.CommandNetworkKeySet:
		fmt.Println("Illegal network key set received (network may be under attack or something)")

	case commandclass.CommandNetworkKeyVerify:
		// @todo implement me
		fmt.Println("Received network key verify")

	case commandclass.CommandSecurityCommandsSupportedGet:
		// @todo implement me maybe
		fmt.Println("Received request for supported security commands")

	case commandclass.CommandSecurityCommandsSupportedReport:
		// @todo implement me maybe
		fmt.Println("Received response for supported security commands")

	case commandclass.CommandSecurityMessageEncapsulation, commandclass.CommandSecurityMessageEncapsulationNonceGet:
		fmt.Println("mother of god...", cmd)
		// 1. decrypt message
		// 2. if it's the first half of a sequenced message, wait for the second half
		// 2.5  if it's an EncapsulationGetNonce, then send a NonceReport back to the sender
		// 3. if it's the second half of a sequenced message, reassemble the payloads
		// 4. emit the payload back to the session layer

	case commandclass.CommandSecurityNonceGet:
		s.handleNonceGet(cmd)

	case commandclass.CommandSecurityNonceReport:
		s.handleNonceReport(cmd)

	case commandclass.CommandSecuritySchemeGet:
		fmt.Println("received request for security scheme")

	case commandclass.CommandSecuritySchemeInherit:
		fmt.Println("received security scheme inherit")

	case commandclass.CommandSecuritySchemeReport:
		fmt.Println("received security scheme report: ", cmd.CommandData)
		if s.secureInclusionMode {
			fmt.Println("in secure inclusion mode")
			s.sendDataSecure(
				cmd.SrcNodeId,
				commandclass.NewSecurityNetworkKeySet(NetworkKey),
				true,
			)
		}

	}
}

func (s *SecurityLayer) sendDataSecure(nodeId uint8, data []byte, inclusionMode bool) error {
	fmt.Println("sending secure data", len(data))
	if len(data) > SecurePayloadMaxSizeAutoRoute {
		fmt.Println("segmenting")
		// we have to split the payload into two frames which are transmitted
		// separately; the second of which must only be transmitted after we receive
		// a NonceGet from the other node and transmit a corresponding NonceReport

		sequenceCounter := s.getNextSequenceCounter(nodeId)
		fmt.Println("seq counter:", sequenceCounter)

		err := s.sendSecurePayload(
			nodeId,
			data[0:SecurePayloadMaxSizeAutoRoute],
			inclusionMode,
			true,
			false,
			sequenceCounter,
		)

		if err != nil {
			fmt.Println("first segment err", err)
			return err
		}

		fmt.Println("sending second segment")

		return s.sendSecurePayload(
			nodeId,
			data[SecurePayloadMaxSizeAutoRoute:],
			inclusionMode,
			true,
			true,
			sequenceCounter,
		)
	} else {
		fmt.Println("not segmenting")
		// we only need to send a single frame
		return s.sendSecurePayload(
			nodeId,
			data,
			inclusionMode,
			false,
			false,
			0,
		)
	}
}

func (s *SecurityLayer) sendSecurePayload(
	nodeId uint8,
	data []byte,
	inclusionMode bool,
	sequenced bool,
	isSecondFrame bool,
	sequenceCounter uint8,
) error {

	var receiverNonce Nonce
	var err error

	if isSecondFrame {
		fmt.Println("sendSecurePayload: waiting for external nonce")
		receiverNonce, err = s.waitForExternalNonce(nodeId)
	} else {
		fmt.Println("sendSecurePayload: requesting external nonce")
		receiverNonce, err = s.getExternalNonce(nodeId)
	}

	if err != nil {
		fmt.Println("sendSecurePayload: nonce err", err)
		return err
	}

	fmt.Printf("nonce from %d: %v\n", nodeId, receiverNonce)

	senderNonce, err := s.internalNonceTable.Generate(InternalNonceTTL)
	if err != nil {
		return err
	}

	var securityByte byte = sequenceCounter & SecuritySequenceCounterMask
	if sequenced {
		securityByte |= SecuritySequenceSequencedFlag

		if isSecondFrame {
			securityByte |= SecuritySequenceSecondFrameFlag
		}
	}

	var encKey, authKey []byte
	if inclusionMode {
		encKey = InclusionEncKey
		authKey = InclusionAuthKey
	} else {
		encKey = NetworkEncKey
		authKey = NetworkAuthKey
	}

	data = append([]byte{securityByte}, data...)

	fmt.Println("input:")
	fmt.Println(hex.Dump(data))

	// full initialization vector = senderNonce + receiverNonce
	iv := append(senderNonce, receiverNonce...)

	fmt.Println("iv:")
	fmt.Println(hex.Dump(iv))

	fmt.Println("enc key:")
	fmt.Println(hex.Dump(encKey))

	encryptedPayload := CryptMessage(data, iv, encKey)

	fmt.Printf("encrypted output (%d):\n", len(encryptedPayload))
	fmt.Println(hex.Dump(encryptedPayload))

	authDataBuf := append(iv, commandclass.CommandSecurityMessageEncapsulation) // @todo CC should be determined by sequencing
	authDataBuf = append(authDataBuf, 1)                                        // sender node
	authDataBuf = append(authDataBuf, nodeId)                                   // receiver node
	authDataBuf = append(authDataBuf, uint8(len(encryptedPayload)))
	authDataBuf = append(authDataBuf, encryptedPayload...)

	fmt.Println("auth iv:")
	fmt.Println(hex.Dump(iv))

	fmt.Println("auth key:")
	fmt.Println(hex.Dump(authKey))

	fmt.Println("auth data:")
	fmt.Println(hex.Dump(authDataBuf))

	authIV := []byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	hmac := CalculateHMAC(authDataBuf, authIV, authKey)

	fmt.Println("hmac:")
	fmt.Println(hex.Dump(hmac))

	encapsulatedMessage := commandclass.NewSecurityMessageEncapsulation(
		senderNonce,
		encryptedPayload,
		hmac,
		receiverNonce[0],
	)

	fmt.Println("encapsulated:")
	fmt.Println(hex.Dump(encapsulatedMessage))

	_, err = s.session.sendDataUnsafe(nodeId, encapsulatedMessage)

	return err
}

func (s *SecurityLayer) getNextSequenceCounter(nodeId uint8) uint8 {
	var counter uint8
	var ok bool

	if counter, ok = s.sequenceCounters[nodeId]; !ok {
		s.sequenceCounters[nodeId] = SecuritySequenceCounterMin
		return SecuritySequenceCounterMin
	}

	if counter+1 > SecuritySequenceCounterMax {
		counter = 0
	} else {
		counter += 1
	}

	s.sequenceCounters[nodeId] = counter

	return counter
}

func (s *SecurityLayer) includeSecureNode(node *Node) error {
	s.secureInclusionMode = true
	s.includingNode = node
	fmt.Printf("secure inclusion: %v\n", s.secureInclusionMode)
	defer s.cleanupSecureInclusionMode()

	done := make(chan bool)

	s.session.sendDataUnsafe(node.NodeId, commandclass.NewSecuritySchemeGet())

	var err error

	select {
	case <-done:
		// @todo check for errors
	case <-time.After(MaxSecureInclusionDuration):
		fmt.Println("secure inclusion timeout")
		err = errors.New("Secure inclusion timed out")
	}

	return err
}

func (s *SecurityLayer) getExternalNonce(nodeId uint8) (Nonce, error) {
	var nonce Nonce
	var err error

	nonce, err = s.externalNonceTable.Get(nodeId)
	if err == nil {
		return nonce, nil
	}

	s.session.sendDataUnsafe(nodeId, commandclass.NewSecurityNonceGet())

	return s.waitForExternalNonce(nodeId)
}

func (s *SecurityLayer) waitForExternalNonce(nodeId uint8) (Nonce, error) {
	var waitChan chan bool
	var ok bool

	// Get the wait channel, creating it if it doesn't exist (note the !ok condition)
	s.waitMapLock.Lock()
	if waitChan, ok = s.waitForNonce[nodeId]; !ok {
		waitChan = make(chan bool)
		s.waitForNonce[nodeId] = waitChan
	}
	s.waitMapLock.Unlock()
	runtime.Gosched()

	fmt.Println("about to start waiting", nodeId, waitChan)
	select {
	case <-waitChan:
		fmt.Println("done waiting for nonce")
	case <-time.After(NonceRequestTimeout):
		fmt.Println("timed out waiting for nonce")
		return nil, errors.New("nonce timeout")
	}

	nonce, err := s.externalNonceTable.Get(nodeId)
	fmt.Println("nonce table:", nonce, err)
	if err == nil {
		return nonce, nil
	}

	return nil, errors.New("Failed to get external nonce")
}

func (s *SecurityLayer) cleanupSecureInclusionMode() {
	fmt.Println("cleaning up secure inclusion mode")
	s.secureInclusionMode = false
	s.includingNode = nil
}

// handleNonceGet will generate a new internal nonce, store it in the security
// layer's internal nonce table, and send a NonceReport to the requester.
// NOTE: The Z-Wave docs are not very clear on this, but the "receiver nonce id"
// is simply the first byte of the nonce (which must be unique among all of the
// active internal nonces)
func (s *SecurityLayer) handleNonceGet(cmd *ApplicationCommandHandlerBridge) {
	nonce, err := s.internalNonceTable.Generate(InternalNonceTTL)

	if err != nil {
		// @todo figure out what to do about this
		panic("unable to generate internal nonce")
	}

	go s.session.sendDataUnsafe(cmd.SrcNodeId, commandclass.NewSecurityNonceReport(nonce))
}

// handleNonceReport stores the received nonce in the security layer's external
// nonce table. Additionally, it sets a timeout on the nonce (after which the
// nonce will be deleted from the nonce table) and notifies any goroutines that
// may be waiting for a nonce from the given node
func (s *SecurityLayer) handleNonceReport(cmd *ApplicationCommandHandlerBridge) {
	cc := commandclass.ParseSecurityNonceReport(cmd.CommandData)
	fmt.Printf("handleNonceReport: received nonce from %d, %v\n", cmd.SrcNodeId, cc.Nonce)
	s.externalNonceTable.Set(cmd.SrcNodeId, cc.Nonce, ExternalNonceTTL)

	// if there is no matching channel in the waitForNonce map, then apparently we
	// either fetched the nonce for no reason, some node just randomly gave us one,
	// or whatever process requested the nonce timed out already. in any case, we've
	// stored the nonce, so it'll be valid for now
	fmt.Println("handleNonceReport: chan", cmd.SrcNodeId)
	if ch, ok := s.waitForNonce[cmd.SrcNodeId]; ok {
		fmt.Println("handleNonceReport: chan ok", cmd.SrcNodeId, ch)

		// perform a non-blocking send. it's possible that some process asked for the
		// nonce, but for whatever reason didn't bother to listen on the channel. in
		// any case, we never want to block here
		select {
		case ch <- true:
			fmt.Println("handleNonceReport: chan wrote", cmd.SrcNodeId)
		default:
			fmt.Println("handleNonceReport: chan not written", cmd.SrcNodeId)
		}

		// closing the channel will unblock anything that is currently listening (in
		// the case of multiple listeners or future listeners). this is especially
		// important if two goroutines are waiting on the channel value, since we only
		// emit one time. this also helps us in the case that some rogue goroutine has
		// a reference to this channel that it hasn't blocked on yet, but will at some
		// point
		close(ch)
		fmt.Println("handleNonceReport: chan closed", cmd.SrcNodeId)

		// delete the channel from the map; a new one will be created when we request
		// another nonce from the node
		delete(s.waitForNonce, cmd.SrcNodeId)
	}
}