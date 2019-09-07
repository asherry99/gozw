// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package firmwareupdatemd

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandGet cc.CommandID = 0x05

func init() {
	gob.Register(Get{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x7A),
		Command:      cc.CommandID(0x05),
		Version:      1,
	}, NewGet)
}

func NewGet() cc.Command {
	return &Get{}
}

// <no value>
type Get struct {
	NumberOfReports byte

	Properties1 struct {
		ReportNumber1 byte

		Zero bool
	}

	ReportNumber2 byte
}

func (cmd Get) CommandClassID() cc.CommandClassID {
	return 0x7A
}

func (cmd Get) CommandID() cc.CommandID {
	return CommandGet
}

func (cmd Get) CommandIDString() string {
	return "FIRMWARE_UPDATE_MD_GET"
}

func (cmd *Get) UnmarshalBinary(data []byte) error {
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

	cmd.NumberOfReports = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.ReportNumber1 = (payload[i] & 0x7F)

	cmd.Properties1.Zero = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.ReportNumber2 = payload[i]
	i++

	return nil
}

func (cmd *Get) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.NumberOfReports)

	{
		var val byte

		val |= (cmd.Properties1.ReportNumber1) & byte(0x7F)

		if cmd.Properties1.Zero {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.ReportNumber2)

	return
}
