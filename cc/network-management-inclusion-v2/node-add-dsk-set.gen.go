// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package networkmanagementinclusionv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandNodeAddDskSet cc.CommandID = 0x14

func init() {
	gob.Register(NodeAddDskSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x34),
		Command:      cc.CommandID(0x14),
		Version:      2,
	}, NewNodeAddDskSet)
}

func NewNodeAddDskSet() cc.Command {
	return &NodeAddDskSet{}
}

// <no value>
type NodeAddDskSet struct {
	SeqNo byte

	Properties1 struct {
		InputDskLength byte

		Accept bool
	}

	InputDsk []byte
}

func (cmd NodeAddDskSet) CommandClassID() cc.CommandClassID {
	return 0x34
}

func (cmd NodeAddDskSet) CommandID() cc.CommandID {
	return CommandNodeAddDskSet
}

func (cmd NodeAddDskSet) CommandIDString() string {
	return "NODE_ADD_DSK_SET"
}

func (cmd *NodeAddDskSet) UnmarshalBinary(data []byte) error {
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

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.InputDskLength = (payload[i] & 0x0F)

	cmd.Properties1.Accept = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return nil // field is optional
	}

	{
		length := (payload[1+2]) & 0x0F
		cmd.InputDsk = payload[i : i+int(length)]
		i += int(length)
	}

	return nil
}

func (cmd *NodeAddDskSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SeqNo)

	{
		var val byte

		val |= (cmd.Properties1.InputDskLength) & byte(0x0F)

		if cmd.Properties1.Accept {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	if cmd.InputDsk != nil && len(cmd.InputDsk) > 0 {
		payload = append(payload, cmd.InputDsk...)
	}

	return
}
