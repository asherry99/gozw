// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package climatecontrolschedule

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandScheduleChangedGet cc.CommandID = 0x04

func init() {
	gob.Register(ScheduleChangedGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x46),
		Command:      cc.CommandID(0x04),
		Version:      1,
	}, NewScheduleChangedGet)
}

func NewScheduleChangedGet() cc.Command {
	return &ScheduleChangedGet{}
}

// <no value>
type ScheduleChangedGet struct {
}

func (cmd ScheduleChangedGet) CommandClassID() cc.CommandClassID {
	return 0x46
}

func (cmd ScheduleChangedGet) CommandID() cc.CommandID {
	return CommandScheduleChangedGet
}

func (cmd ScheduleChangedGet) CommandIDString() string {
	return "SCHEDULE_CHANGED_GET"
}

func (cmd *ScheduleChangedGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *ScheduleChangedGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}