// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package firmwareupdatemdv4

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandFirmwareMdGet cc.CommandID = 0x01

func init() {
	gob.Register(FirmwareMdGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x7A),
		Command:      cc.CommandID(0x01),
		Version:      4,
	}, NewFirmwareMdGet)
}

func NewFirmwareMdGet() cc.Command {
	return &FirmwareMdGet{}
}

// <no value>
type FirmwareMdGet struct {
}

func (cmd FirmwareMdGet) CommandClassID() cc.CommandClassID {
	return 0x7A
}

func (cmd FirmwareMdGet) CommandID() cc.CommandID {
	return CommandFirmwareMdGet
}

func (cmd FirmwareMdGet) CommandIDString() string {
	return "FIRMWARE_MD_GET"
}

func (cmd *FirmwareMdGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *FirmwareMdGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}