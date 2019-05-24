// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package notificationv4

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandEventSupportedReport cc.CommandID = 0x02

func init() {
	gob.Register(EventSupportedReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x71),
		Command:      cc.CommandID(0x02),
		Version:      4,
	}, NewEventSupportedReport)
}

func NewEventSupportedReport() cc.Command {
	return &EventSupportedReport{}
}

// <no value>
type EventSupportedReport struct {
	NotificationType byte

	Properties1 struct {
		NumberOfBitMasks byte
	}

	BitMask []byte
}

func (cmd EventSupportedReport) CommandClassID() cc.CommandClassID {
	return 0x71
}

func (cmd EventSupportedReport) CommandID() cc.CommandID {
	return CommandEventSupportedReport
}

func (cmd EventSupportedReport) CommandIDString() string {
	return "EVENT_SUPPORTED_REPORT"
}

func (cmd *EventSupportedReport) UnmarshalBinary(data []byte) error {
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

	cmd.NotificationType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.NumberOfBitMasks = (payload[i] & 0x1F)

	i += 1

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.BitMask = payload[i:]

	return nil
}

func (cmd *EventSupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.NotificationType)

	{
		var val byte

		val |= (cmd.Properties1.NumberOfBitMasks) & byte(0x1F)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.BitMask...)

	return
}
