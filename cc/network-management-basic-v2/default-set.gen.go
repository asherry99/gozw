// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package networkmanagementbasicv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandDefaultSet cc.CommandID = 0x06

func init() {
	gob.Register(DefaultSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x4D),
		Command:      cc.CommandID(0x06),
		Version:      2,
	}, NewDefaultSet)
}

func NewDefaultSet() cc.Command {
	return &DefaultSet{}
}

// <no value>
type DefaultSet struct {
	SeqNo byte
}

func (cmd DefaultSet) CommandClassID() cc.CommandClassID {
	return 0x4D
}

func (cmd DefaultSet) CommandID() cc.CommandID {
	return CommandDefaultSet
}

func (cmd DefaultSet) CommandIDString() string {
	return "DEFAULT_SET"
}

func (cmd *DefaultSet) UnmarshalBinary(data []byte) error {
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

func (cmd *DefaultSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SeqNo)

	return
}
