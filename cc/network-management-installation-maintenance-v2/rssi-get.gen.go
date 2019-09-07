// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package networkmanagementinstallationmaintenancev2

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandRssiGet cc.CommandID = 0x07

func init() {
	gob.Register(RssiGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x67),
		Command:      cc.CommandID(0x07),
		Version:      2,
	}, NewRssiGet)
}

func NewRssiGet() cc.Command {
	return &RssiGet{}
}

// <no value>
type RssiGet struct {
}

func (cmd RssiGet) CommandClassID() cc.CommandClassID {
	return 0x67
}

func (cmd RssiGet) CommandID() cc.CommandID {
	return CommandRssiGet
}

func (cmd RssiGet) CommandIDString() string {
	return "RSSI_GET"
}

func (cmd *RssiGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *RssiGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
