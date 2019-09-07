// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package zipnaming

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandLocationGet cc.CommandID = 0x05

func init() {
	gob.Register(LocationGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x68),
		Command:      cc.CommandID(0x05),
		Version:      1,
	}, NewLocationGet)
}

func NewLocationGet() cc.Command {
	return &LocationGet{}
}

// <no value>
type LocationGet struct {
}

func (cmd LocationGet) CommandClassID() cc.CommandClassID {
	return 0x68
}

func (cmd LocationGet) CommandID() cc.CommandID {
	return CommandLocationGet
}

func (cmd LocationGet) CommandIDString() string {
	return "ZIP_NAMING_LOCATION_GET"
}

func (cmd *LocationGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *LocationGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
