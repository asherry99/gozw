// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package security2

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandKexGet cc.CommandID = 0x04

func init() {
	gob.Register(KexGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x9F),
		Command:      cc.CommandID(0x04),
		Version:      1,
	}, NewKexGet)
}

func NewKexGet() cc.Command {
	return &KexGet{}
}

// <no value>
type KexGet struct {
}

func (cmd KexGet) CommandClassID() cc.CommandClassID {
	return 0x9F
}

func (cmd KexGet) CommandID() cc.CommandID {
	return CommandKexGet
}

func (cmd KexGet) CommandIDString() string {
	return "KEX_GET"
}

func (cmd *KexGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *KexGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
