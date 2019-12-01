// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package scheduleentrylockv3

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandWeekDayGet cc.CommandID = 0x04

func init() {
	gob.Register(WeekDayGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x4E),
		Command:      cc.CommandID(0x04),
		Version:      3,
	}, NewWeekDayGet)
}

func NewWeekDayGet() cc.Command {
	return &WeekDayGet{}
}

// <no value>
type WeekDayGet struct {
	UserIdentifier byte

	ScheduleSlotId byte
}

func (cmd WeekDayGet) CommandClassID() cc.CommandClassID {
	return 0x4E
}

func (cmd WeekDayGet) CommandID() cc.CommandID {
	return CommandWeekDayGet
}

func (cmd WeekDayGet) CommandIDString() string {
	return "SCHEDULE_ENTRY_LOCK_WEEK_DAY_GET"
}

func (cmd *WeekDayGet) UnmarshalBinary(data []byte) error {
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

	cmd.UserIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ScheduleSlotId = payload[i]
	i++

	return nil
}

func (cmd *WeekDayGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.UserIdentifier)

	payload = append(payload, cmd.ScheduleSlotId)

	return
}