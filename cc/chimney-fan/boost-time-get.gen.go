// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package chimneyfan

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandBoostTimeGet cc.CommandID = 0x11

func init() {
	gob.Register(BoostTimeGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x2A),
		Command:      cc.CommandID(0x11),
		Version:      1,
	}, NewBoostTimeGet)
}

func NewBoostTimeGet() cc.Command {
	return &BoostTimeGet{}
}

// <no value>
type BoostTimeGet struct {
}

func (cmd BoostTimeGet) CommandClassID() cc.CommandClassID {
	return 0x2A
}

func (cmd BoostTimeGet) CommandID() cc.CommandID {
	return CommandBoostTimeGet
}

func (cmd BoostTimeGet) CommandIDString() string {
	return "CHIMNEY_FAN_BOOST_TIME_GET"
}

func (cmd *BoostTimeGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *BoostTimeGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
