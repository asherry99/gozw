// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package zipnd

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandZipNodeSolicitation cc.CommandID = 0x03

func init() {
	gob.Register(ZipNodeSolicitation{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x58),
		Command:      cc.CommandID(0x03),
		Version:      1,
	}, NewZipNodeSolicitation)
}

func NewZipNodeSolicitation() cc.Command {
	return &ZipNodeSolicitation{}
}

// <no value>
type ZipNodeSolicitation struct {
	NodeId byte

	Ipv6Address []byte
}

func (cmd ZipNodeSolicitation) CommandClassID() cc.CommandClassID {
	return 0x58
}

func (cmd ZipNodeSolicitation) CommandID() cc.CommandID {
	return CommandZipNodeSolicitation
}

func (cmd ZipNodeSolicitation) CommandIDString() string {
	return "ZIP_NODE_SOLICITATION"
}

func (cmd *ZipNodeSolicitation) UnmarshalBinary(data []byte) error {
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

	i++ // skipping over reserved bit

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NodeId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Ipv6Address = payload[i : i+16]

	i += 16

	return nil
}

func (cmd *ZipNodeSolicitation) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.NodeId)

	if paramLen := len(cmd.Ipv6Address); paramLen > 16 {
		return nil, errors.New("Length overflow in array parameter Ipv6Address")
	}

	payload = append(payload, cmd.Ipv6Address...)

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}
