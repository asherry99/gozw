// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package scheduleentrylockv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandYearDaySet cc.CommandID = 0x06

func init() {
	gob.Register(YearDaySet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x4E),
		Command:      cc.CommandID(0x06),
		Version:      2,
	}, NewYearDaySet)
}

func NewYearDaySet() cc.Command {
	return &YearDaySet{}
}

// <no value>
type YearDaySet struct {
	SetAction byte

	UserIdentifier byte

	ScheduleSlotId byte

	StartYear byte

	StartMonth byte

	StartDay byte

	StartHour byte

	StartMinute byte

	StopYear byte

	StopMonth byte

	StopDay byte

	StopHour byte

	StopMinute byte
}

func (cmd YearDaySet) CommandClassID() cc.CommandClassID {
	return 0x4E
}

func (cmd YearDaySet) CommandID() cc.CommandID {
	return CommandYearDaySet
}

func (cmd YearDaySet) CommandIDString() string {
	return "SCHEDULE_ENTRY_LOCK_YEAR_DAY_SET"
}

func (cmd *YearDaySet) UnmarshalBinary(data []byte) error {
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

	cmd.SetAction = payload[i]
	i++

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

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartYear = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartMonth = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartDay = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartHour = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartMinute = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopYear = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopMonth = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopDay = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopHour = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopMinute = payload[i]
	i++

	return nil
}

func (cmd *YearDaySet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SetAction)

	payload = append(payload, cmd.UserIdentifier)

	payload = append(payload, cmd.ScheduleSlotId)

	payload = append(payload, cmd.StartYear)

	payload = append(payload, cmd.StartMonth)

	payload = append(payload, cmd.StartDay)

	payload = append(payload, cmd.StartHour)

	payload = append(payload, cmd.StartMinute)

	payload = append(payload, cmd.StopYear)

	payload = append(payload, cmd.StopMonth)

	payload = append(payload, cmd.StopDay)

	payload = append(payload, cmd.StopHour)

	payload = append(payload, cmd.StopMinute)

	return
}
