// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package schedulev3

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandStateSet cc.CommandID = 0x07

func init() {
	gob.Register(StateSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x53),
		Command:      cc.CommandID(0x07),
		Version:      3,
	}, NewStateSet)
}

func NewStateSet() cc.Command {
	return &StateSet{}
}

// <no value>
type StateSet struct {
	ScheduleId byte

	ScheduleState byte

	ScheduleIdBlock byte
}

func (cmd StateSet) CommandClassID() cc.CommandClassID {
	return 0x53
}

func (cmd StateSet) CommandID() cc.CommandID {
	return CommandStateSet
}

func (cmd StateSet) CommandIDString() string {
	return "SCHEDULE_STATE_SET"
}

func (cmd *StateSet) UnmarshalBinary(data []byte) error {
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

	cmd.ScheduleId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ScheduleState = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ScheduleIdBlock = payload[i]
	i++

	return nil
}

func (cmd *StateSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.ScheduleId)

	payload = append(payload, cmd.ScheduleState)

	payload = append(payload, cmd.ScheduleIdBlock)

	return
}
