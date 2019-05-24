// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package notificationv3

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandSupportedReport cc.CommandID = 0x08

func init() {
	gob.Register(SupportedReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x71),
		Command:      cc.CommandID(0x08),
		Version:      3,
	}, NewSupportedReport)
}

func NewSupportedReport() cc.Command {
	return &SupportedReport{}
}

// <no value>
type SupportedReport struct {
	Properties1 struct {
		NumberOfBitMasks byte

		V1Alarm bool
	}

	BitMask []byte
}

func (cmd SupportedReport) CommandClassID() cc.CommandClassID {
	return 0x71
}

func (cmd SupportedReport) CommandID() cc.CommandID {
	return CommandSupportedReport
}

func (cmd SupportedReport) CommandIDString() string {
	return "NOTIFICATION_SUPPORTED_REPORT"
}

func (cmd *SupportedReport) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.NumberOfBitMasks = (payload[i] & 0x1F)

	cmd.Properties1.V1Alarm = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.BitMask = payload[i:]

	return nil
}

func (cmd *SupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.NumberOfBitMasks) & byte(0x1F)

		if cmd.Properties1.V1Alarm {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.BitMask...)

	return
}
