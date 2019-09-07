// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package scheduleentrylockv3

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandScheduleEntryTypeSupportedGet cc.CommandID = 0x09

func init() {
	gob.Register(ScheduleEntryTypeSupportedGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x4E),
		Command:      cc.CommandID(0x09),
		Version:      3,
	}, NewScheduleEntryTypeSupportedGet)
}

func NewScheduleEntryTypeSupportedGet() cc.Command {
	return &ScheduleEntryTypeSupportedGet{}
}

// <no value>
type ScheduleEntryTypeSupportedGet struct {
}

func (cmd ScheduleEntryTypeSupportedGet) CommandClassID() cc.CommandClassID {
	return 0x4E
}

func (cmd ScheduleEntryTypeSupportedGet) CommandID() cc.CommandID {
	return CommandScheduleEntryTypeSupportedGet
}

func (cmd ScheduleEntryTypeSupportedGet) CommandIDString() string {
	return "SCHEDULE_ENTRY_TYPE_SUPPORTED_GET"
}

func (cmd *ScheduleEntryTypeSupportedGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *ScheduleEntryTypeSupportedGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
