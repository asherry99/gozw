// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package scheduleentrylockv3

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandYearDayGet cc.CommandID = 0x07

func init() {
	gob.Register(YearDayGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x4E),
		Command:      cc.CommandID(0x07),
		Version:      3,
	}, NewYearDayGet)
}

func NewYearDayGet() cc.Command {
	return &YearDayGet{}
}

// <no value>
type YearDayGet struct {
	UserIdentifier byte

	ScheduleSlotId byte
}

func (cmd YearDayGet) CommandClassID() cc.CommandClassID {
	return 0x4E
}

func (cmd YearDayGet) CommandID() cc.CommandID {
	return CommandYearDayGet
}

func (cmd YearDayGet) CommandIDString() string {
	return "SCHEDULE_ENTRY_LOCK_YEAR_DAY_GET"
}

func (cmd *YearDayGet) UnmarshalBinary(data []byte) error {
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

func (cmd *YearDayGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.UserIdentifier)

	payload = append(payload, cmd.ScheduleSlotId)

	return
}
