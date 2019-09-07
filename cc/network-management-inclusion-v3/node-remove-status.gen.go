// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package networkmanagementinclusionv3

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandNodeRemoveStatus cc.CommandID = 0x04

func init() {
	gob.Register(NodeRemoveStatus{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x34),
		Command:      cc.CommandID(0x04),
		Version:      3,
	}, NewNodeRemoveStatus)
}

func NewNodeRemoveStatus() cc.Command {
	return &NodeRemoveStatus{}
}

// <no value>
type NodeRemoveStatus struct {
	SeqNo byte

	Status byte

	Nodeid byte
}

func (cmd NodeRemoveStatus) CommandClassID() cc.CommandClassID {
	return 0x34
}

func (cmd NodeRemoveStatus) CommandID() cc.CommandID {
	return CommandNodeRemoveStatus
}

func (cmd NodeRemoveStatus) CommandIDString() string {
	return "NODE_REMOVE_STATUS"
}

func (cmd *NodeRemoveStatus) UnmarshalBinary(data []byte) error {
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

	cmd.Status = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Nodeid = payload[i]
	i++

	return nil
}

func (cmd *NodeRemoveStatus) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SeqNo)

	payload = append(payload, cmd.Status)

	payload = append(payload, cmd.Nodeid)

	return
}
