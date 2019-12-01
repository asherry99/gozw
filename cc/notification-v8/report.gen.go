// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package notificationv8

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandReport cc.CommandID = 0x05

func init() {
	gob.Register(Report{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x71),
		Command:      cc.CommandID(0x05),
		Version:      8,
	}, NewReport)
}

func NewReport() cc.Command {
	return &Report{}
}

// <no value>
type Report struct {
	V1AlarmType byte

	V1AlarmLevel byte

	NotificationStatus byte

	NotificationType byte

	Event byte

	Properties1 struct {
		EventParametersLength byte

		Sequence bool
	}

	EventParameter []byte

	SequenceNumber byte
}

func (cmd Report) CommandClassID() cc.CommandClassID {
	return 0x71
}

func (cmd Report) CommandID() cc.CommandID {
	return CommandReport
}

func (cmd Report) CommandIDString() string {
	return "NOTIFICATION_REPORT"
}

func (cmd *Report) UnmarshalBinary(data []byte) error {
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

	cmd.V1AlarmType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.V1AlarmLevel = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	i++ // skipping over reserved bit

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NotificationStatus = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NotificationType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Event = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.EventParametersLength = (payload[i] & 0x1F)

	cmd.Properties1.Sequence = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return nil // field is optional
	}

	{
		length := (payload[6+2]) & 0x1F
		cmd.EventParameter = payload[i : i+int(length)]
		i += int(length)
	}

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.SequenceNumber = payload[i]
	i++

	return nil
}

func (cmd *Report) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.V1AlarmType)

	payload = append(payload, cmd.V1AlarmLevel)

	payload = append(payload, cmd.NotificationStatus)

	payload = append(payload, cmd.NotificationType)

	payload = append(payload, cmd.Event)

	{
		var val byte

		val |= (cmd.Properties1.EventParametersLength) & byte(0x1F)

		if cmd.Properties1.Sequence {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	if cmd.EventParameter != nil && len(cmd.EventParameter) > 0 {
		payload = append(payload, cmd.EventParameter...)
	}

	payload = append(payload, cmd.SequenceNumber)

	return
}