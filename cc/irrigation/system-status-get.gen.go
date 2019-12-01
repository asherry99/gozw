// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package irrigation

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandSystemStatusGet cc.CommandID = 0x03

func init() {
	gob.Register(SystemStatusGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x6B),
		Command:      cc.CommandID(0x03),
		Version:      1,
	}, NewSystemStatusGet)
}

func NewSystemStatusGet() cc.Command {
	return &SystemStatusGet{}
}

// <no value>
type SystemStatusGet struct {
}

func (cmd SystemStatusGet) CommandClassID() cc.CommandClassID {
	return 0x6B
}

func (cmd SystemStatusGet) CommandID() cc.CommandID {
	return CommandSystemStatusGet
}

func (cmd SystemStatusGet) CommandIDString() string {
	return "IRRIGATION_SYSTEM_STATUS_GET"
}

func (cmd *SystemStatusGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *SystemStatusGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}