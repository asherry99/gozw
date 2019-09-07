// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package schedulev4

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandSupportedGet cc.CommandID = 0x01

func init() {
	gob.Register(SupportedGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x53),
		Command:      cc.CommandID(0x01),
		Version:      4,
	}, NewSupportedGet)
}

func NewSupportedGet() cc.Command {
	return &SupportedGet{}
}

// <no value>
type SupportedGet struct {
	ScheduleIdBlock byte
}

func (cmd SupportedGet) CommandClassID() cc.CommandClassID {
	return 0x53
}

func (cmd SupportedGet) CommandID() cc.CommandID {
	return CommandSupportedGet
}

func (cmd SupportedGet) CommandIDString() string {
	return "SCHEDULE_SUPPORTED_GET"
}

func (cmd *SupportedGet) UnmarshalBinary(data []byte) error {
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

	cmd.ScheduleIdBlock = payload[i]
	i++

	return nil
}

func (cmd *SupportedGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.ScheduleIdBlock)

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}
