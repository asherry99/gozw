// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package chimneyfan

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandAlarmStatusSet cc.CommandID = 0x22

func init() {
	gob.Register(AlarmStatusSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x2A),
		Command:      cc.CommandID(0x22),
		Version:      1,
	}, NewAlarmStatusSet)
}

func NewAlarmStatusSet() cc.Command {
	return &AlarmStatusSet{}
}

// <no value>
type AlarmStatusSet struct {
	Message struct {
		NotUsed2 byte

		NotUsed1 bool

		AcknowledgeExternalAlarm bool

		AcknowledgeSensorError bool

		AcknowledgeAlarmTemperatureExceeded bool
	}
}

func (cmd AlarmStatusSet) CommandClassID() cc.CommandClassID {
	return 0x2A
}

func (cmd AlarmStatusSet) CommandID() cc.CommandID {
	return CommandAlarmStatusSet
}

func (cmd AlarmStatusSet) CommandIDString() string {
	return "CHIMNEY_FAN_ALARM_STATUS_SET"
}

func (cmd *AlarmStatusSet) UnmarshalBinary(data []byte) error {
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

	cmd.Message.NotUsed2 = (payload[i] & 0xF0) >> 4

	cmd.Message.NotUsed1 = payload[i]&0x01 == 0x01

	cmd.Message.AcknowledgeExternalAlarm = payload[i]&0x02 == 0x02

	cmd.Message.AcknowledgeSensorError = payload[i]&0x04 == 0x04

	cmd.Message.AcknowledgeAlarmTemperatureExceeded = payload[i]&0x08 == 0x08

	i += 1

	return nil
}

func (cmd *AlarmStatusSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Message.NotUsed2 << byte(4)) & byte(0xF0)

		if cmd.Message.NotUsed1 {
			val |= byte(0x01) // flip bits on
		} else {
			val &= ^byte(0x01) // flip bits off
		}

		if cmd.Message.AcknowledgeExternalAlarm {
			val |= byte(0x02) // flip bits on
		} else {
			val &= ^byte(0x02) // flip bits off
		}

		if cmd.Message.AcknowledgeSensorError {
			val |= byte(0x04) // flip bits on
		} else {
			val &= ^byte(0x04) // flip bits off
		}

		if cmd.Message.AcknowledgeAlarmTemperatureExceeded {
			val |= byte(0x08) // flip bits on
		} else {
			val &= ^byte(0x08) // flip bits off
		}

		payload = append(payload, val)
	}

	return
}
