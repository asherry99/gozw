// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package networkmanagementproxyv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandNodeListGet cc.CommandID = 0x01

func init() {
	gob.Register(NodeListGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x52),
		Command:      cc.CommandID(0x01),
		Version:      2,
	}, NewNodeListGet)
}

func NewNodeListGet() cc.Command {
	return &NodeListGet{}
}

// <no value>
type NodeListGet struct {
	SeqNo byte
}

func (cmd NodeListGet) CommandClassID() cc.CommandClassID {
	return 0x52
}

func (cmd NodeListGet) CommandID() cc.CommandID {
	return CommandNodeListGet
}

func (cmd NodeListGet) CommandIDString() string {
	return "NODE_LIST_GET"
}

func (cmd *NodeListGet) UnmarshalBinary(data []byte) error {
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

func (cmd *NodeListGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SeqNo)

	return
}
