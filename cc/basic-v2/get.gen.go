// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package basicv2

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandGet cc.CommandID = 0x02

func init() {
	gob.Register(Get{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x20),
		Command:      cc.CommandID(0x02),
		Version:      2,
	}, NewGet)
}

func NewGet() cc.Command {
	return &Get{}
}

// <no value>
type Get struct {
}

func (cmd Get) CommandClassID() cc.CommandClassID {
	return 0x20
}

func (cmd Get) CommandID() cc.CommandID {
	return CommandGet
}

func (cmd Get) CommandIDString() string {
	return "BASIC_GET"
}

func (cmd *Get) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *Get) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}
