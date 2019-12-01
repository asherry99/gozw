// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package climatecontrolschedule

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandScheduleOverrideSet cc.CommandID = 0x06

func init() {
	gob.Register(ScheduleOverrideSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x46),
		Command:      cc.CommandID(0x06),
		Version:      1,
	}, NewScheduleOverrideSet)
}

func NewScheduleOverrideSet() cc.Command {
	return &ScheduleOverrideSet{}
}

// <no value>
type ScheduleOverrideSet struct {
	Properties1 struct {
		OverrideType byte
	}

	OverrideState byte
}

func (cmd ScheduleOverrideSet) CommandClassID() cc.CommandClassID {
	return 0x46
}

func (cmd ScheduleOverrideSet) CommandID() cc.CommandID {
	return CommandScheduleOverrideSet
}

func (cmd ScheduleOverrideSet) CommandIDString() string {
	return "SCHEDULE_OVERRIDE_SET"
}

func (cmd *ScheduleOverrideSet) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.OverrideType = (payload[i] & 0x03)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.OverrideState = payload[i]
	i++

	return nil
}

func (cmd *ScheduleOverrideSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.OverrideType) & byte(0x03)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.OverrideState)

	return
}