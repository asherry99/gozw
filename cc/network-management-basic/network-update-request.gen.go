// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package networkmanagementbasic

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandNetworkUpdateRequest cc.CommandID = 0x03

func init() {
	gob.Register(NetworkUpdateRequest{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x4D),
		Command:      cc.CommandID(0x03),
		Version:      1,
	}, NewNetworkUpdateRequest)
}

func NewNetworkUpdateRequest() cc.Command {
	return &NetworkUpdateRequest{}
}

// <no value>
type NetworkUpdateRequest struct {
	SeqNo byte
}

func (cmd NetworkUpdateRequest) CommandClassID() cc.CommandClassID {
	return 0x4D
}

func (cmd NetworkUpdateRequest) CommandID() cc.CommandID {
	return CommandNetworkUpdateRequest
}

func (cmd NetworkUpdateRequest) CommandIDString() string {
	return "NETWORK_UPDATE_REQUEST"
}

func (cmd *NetworkUpdateRequest) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SeqNo = payload[i]
	i++

	return nil
}

func (cmd *NetworkUpdateRequest) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SeqNo)

	return
}
