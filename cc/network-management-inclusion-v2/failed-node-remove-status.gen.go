// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package networkmanagementinclusionv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandFailedNodeRemoveStatus cc.CommandID = 0x08

func init() {
	gob.Register(FailedNodeRemoveStatus{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x34),
		Command:      cc.CommandID(0x08),
		Version:      2,
	}, NewFailedNodeRemoveStatus)
}

func NewFailedNodeRemoveStatus() cc.Command {
	return &FailedNodeRemoveStatus{}
}

// <no value>
type FailedNodeRemoveStatus struct {
	SeqNo byte

	Status byte

	NodeId byte
}

func (cmd FailedNodeRemoveStatus) CommandClassID() cc.CommandClassID {
	return 0x34
}

func (cmd FailedNodeRemoveStatus) CommandID() cc.CommandID {
	return CommandFailedNodeRemoveStatus
}

func (cmd FailedNodeRemoveStatus) CommandIDString() string {
	return "FAILED_NODE_REMOVE_STATUS"
}

func (cmd *FailedNodeRemoveStatus) UnmarshalBinary(data []byte) error {
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

	cmd.NodeId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Status = payload[i]
	i++

	return nil
}

func (cmd *FailedNodeRemoveStatus) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SeqNo)

	payload = append(payload, cmd.NodeId)

	payload = append(payload, cmd.Status)

	return
}
