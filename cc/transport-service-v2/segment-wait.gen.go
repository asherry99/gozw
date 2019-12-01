// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package transportservicev2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandSegmentWait cc.CommandID = 0xF0

func init() {
	gob.Register(SegmentWait{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x55),
		Command:      cc.CommandID(0xF0),
		Version:      2,
	}, NewSegmentWait)
}

func NewSegmentWait() cc.Command {
	return &SegmentWait{}
}

// <no value>
type SegmentWait struct {
	Properties1 struct {
	}

	PendingFragments byte
}

func (cmd SegmentWait) CommandClassID() cc.CommandClassID {
	return 0x55
}

func (cmd SegmentWait) CommandID() cc.CommandID {
	return CommandSegmentWait
}

func (cmd SegmentWait) CommandIDString() string {
	return "COMMAND_SEGMENT_WAIT"
}

func (cmd *SegmentWait) UnmarshalBinary(data []byte) error {
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

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.PendingFragments = payload[i]
	i++

	return nil
}

func (cmd *SegmentWait) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		payload = append(payload, val)
	}

	payload = append(payload, cmd.PendingFragments)

	return
}