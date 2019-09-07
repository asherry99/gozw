// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package networkmanagementinclusionv3

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandNodeNeighborUpdateStatus cc.CommandID = 0x0C

func init() {
	gob.Register(NodeNeighborUpdateStatus{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x34),
		Command:      cc.CommandID(0x0C),
		Version:      3,
	}, NewNodeNeighborUpdateStatus)
}

func NewNodeNeighborUpdateStatus() cc.Command {
	return &NodeNeighborUpdateStatus{}
}

// <no value>
type NodeNeighborUpdateStatus struct {
	SeqNo byte

	Status byte
}

func (cmd NodeNeighborUpdateStatus) CommandClassID() cc.CommandClassID {
	return 0x34
}

func (cmd NodeNeighborUpdateStatus) CommandID() cc.CommandID {
	return CommandNodeNeighborUpdateStatus
}

func (cmd NodeNeighborUpdateStatus) CommandIDString() string {
	return "NODE_NEIGHBOR_UPDATE_STATUS"
}

func (cmd *NodeNeighborUpdateStatus) UnmarshalBinary(data []byte) error {
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

	return nil
}

func (cmd *NodeNeighborUpdateStatus) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SeqNo)

	payload = append(payload, cmd.Status)

	return
}
