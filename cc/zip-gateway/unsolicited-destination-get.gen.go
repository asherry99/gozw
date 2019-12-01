// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package zipgateway

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandUnsolicitedDestinationGet cc.CommandID = 0x09

func init() {
	gob.Register(UnsolicitedDestinationGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x5F),
		Command:      cc.CommandID(0x09),
		Version:      1,
	}, NewUnsolicitedDestinationGet)
}

func NewUnsolicitedDestinationGet() cc.Command {
	return &UnsolicitedDestinationGet{}
}

// <no value>
type UnsolicitedDestinationGet struct {
}

func (cmd UnsolicitedDestinationGet) CommandClassID() cc.CommandClassID {
	return 0x5F
}

func (cmd UnsolicitedDestinationGet) CommandID() cc.CommandID {
	return CommandUnsolicitedDestinationGet
}

func (cmd UnsolicitedDestinationGet) CommandIDString() string {
	return "UNSOLICITED_DESTINATION_GET"
}

func (cmd *UnsolicitedDestinationGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *UnsolicitedDestinationGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}