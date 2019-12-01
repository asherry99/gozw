// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package hrvcontrol

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandModeSupportedGet cc.CommandID = 0x0A

func init() {
	gob.Register(ModeSupportedGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x39),
		Command:      cc.CommandID(0x0A),
		Version:      1,
	}, NewModeSupportedGet)
}

func NewModeSupportedGet() cc.Command {
	return &ModeSupportedGet{}
}

// <no value>
type ModeSupportedGet struct {
}

func (cmd ModeSupportedGet) CommandClassID() cc.CommandClassID {
	return 0x39
}

func (cmd ModeSupportedGet) CommandID() cc.CommandID {
	return CommandModeSupportedGet
}

func (cmd ModeSupportedGet) CommandIDString() string {
	return "HRV_CONTROL_MODE_SUPPORTED_GET"
}

func (cmd *ModeSupportedGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *ModeSupportedGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}