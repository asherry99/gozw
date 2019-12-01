// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package chimneyfan

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandAlarmLogReport cc.CommandID = 0x21

func init() {
	gob.Register(AlarmLogReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x2A),
		Command:      cc.CommandID(0x21),
		Version:      1,
	}, NewAlarmLogReport)
}

func NewAlarmLogReport() cc.Command {
	return &AlarmLogReport{}
}

// <no value>
type AlarmLogReport struct {
	AlarmEvent1 struct {
		ExternalAlarm1 bool

		SensorError1 bool

		AlarmTemperatureExceeded1 bool

		AlarmStillActive1 bool
	}

	AlarmEvent2 struct {
		ExternalAlarm2 bool

		SensorError2 bool

		AlarmTemperatureExceeded2 bool

		AlarmStillActive2 bool
	}

	AlarmEvent3 struct {
		ExternalAlarm3 bool

		SensorError3 bool

		AlarmTemperatureExceeded3 bool

		AlarmStillActive3 bool
	}

	AlarmEvent4 struct {
		ExternalAlarm4 bool

		SensorError4 bool

		AlarmTemperatureExceeded4 bool

		AlarmStillActive4 bool
	}

	AlarmEvent5 struct {
		ExternalAlarm5 bool

		SensorError5 bool

		AlarmTemperatureExceeded5 bool

		AlarmStillActive5 bool
	}
}

func (cmd AlarmLogReport) CommandClassID() cc.CommandClassID {
	return 0x2A
}

func (cmd AlarmLogReport) CommandID() cc.CommandID {
	return CommandAlarmLogReport
}

func (cmd AlarmLogReport) CommandIDString() string {
	return "CHIMNEY_FAN_ALARM_LOG_REPORT"
}

func (cmd *AlarmLogReport) UnmarshalBinary(data []byte) error {
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

	cmd.AlarmEvent1.ExternalAlarm1 = payload[i]&0x02 == 0x02

	cmd.AlarmEvent1.SensorError1 = payload[i]&0x04 == 0x04

	cmd.AlarmEvent1.AlarmTemperatureExceeded1 = payload[i]&0x08 == 0x08

	cmd.AlarmEvent1.AlarmStillActive1 = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.AlarmEvent2.ExternalAlarm2 = payload[i]&0x02 == 0x02

	cmd.AlarmEvent2.SensorError2 = payload[i]&0x04 == 0x04

	cmd.AlarmEvent2.AlarmTemperatureExceeded2 = payload[i]&0x08 == 0x08

	cmd.AlarmEvent2.AlarmStillActive2 = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.AlarmEvent3.ExternalAlarm3 = payload[i]&0x02 == 0x02

	cmd.AlarmEvent3.SensorError3 = payload[i]&0x04 == 0x04

	cmd.AlarmEvent3.AlarmTemperatureExceeded3 = payload[i]&0x08 == 0x08

	cmd.AlarmEvent3.AlarmStillActive3 = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.AlarmEvent4.ExternalAlarm4 = payload[i]&0x02 == 0x02

	cmd.AlarmEvent4.SensorError4 = payload[i]&0x04 == 0x04

	cmd.AlarmEvent4.AlarmTemperatureExceeded4 = payload[i]&0x08 == 0x08

	cmd.AlarmEvent4.AlarmStillActive4 = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.AlarmEvent5.ExternalAlarm5 = payload[i]&0x02 == 0x02

	cmd.AlarmEvent5.SensorError5 = payload[i]&0x04 == 0x04

	cmd.AlarmEvent5.AlarmTemperatureExceeded5 = payload[i]&0x08 == 0x08

	cmd.AlarmEvent5.AlarmStillActive5 = payload[i]&0x80 == 0x80

	i += 1

	return nil
}

func (cmd *AlarmLogReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		if cmd.AlarmEvent1.ExternalAlarm1 {
			val |= byte(0x02) // flip bits on
		} else {
			val &= ^byte(0x02) // flip bits off
		}

		if cmd.AlarmEvent1.SensorError1 {
			val |= byte(0x04) // flip bits on
		} else {
			val &= ^byte(0x04) // flip bits off
		}

		if cmd.AlarmEvent1.AlarmTemperatureExceeded1 {
			val |= byte(0x08) // flip bits on
		} else {
			val &= ^byte(0x08) // flip bits off
		}

		if cmd.AlarmEvent1.AlarmStillActive1 {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	{
		var val byte

		if cmd.AlarmEvent2.ExternalAlarm2 {
			val |= byte(0x02) // flip bits on
		} else {
			val &= ^byte(0x02) // flip bits off
		}

		if cmd.AlarmEvent2.SensorError2 {
			val |= byte(0x04) // flip bits on
		} else {
			val &= ^byte(0x04) // flip bits off
		}

		if cmd.AlarmEvent2.AlarmTemperatureExceeded2 {
			val |= byte(0x08) // flip bits on
		} else {
			val &= ^byte(0x08) // flip bits off
		}

		if cmd.AlarmEvent2.AlarmStillActive2 {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	{
		var val byte

		if cmd.AlarmEvent3.ExternalAlarm3 {
			val |= byte(0x02) // flip bits on
		} else {
			val &= ^byte(0x02) // flip bits off
		}

		if cmd.AlarmEvent3.SensorError3 {
			val |= byte(0x04) // flip bits on
		} else {
			val &= ^byte(0x04) // flip bits off
		}

		if cmd.AlarmEvent3.AlarmTemperatureExceeded3 {
			val |= byte(0x08) // flip bits on
		} else {
			val &= ^byte(0x08) // flip bits off
		}

		if cmd.AlarmEvent3.AlarmStillActive3 {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	{
		var val byte

		if cmd.AlarmEvent4.ExternalAlarm4 {
			val |= byte(0x02) // flip bits on
		} else {
			val &= ^byte(0x02) // flip bits off
		}

		if cmd.AlarmEvent4.SensorError4 {
			val |= byte(0x04) // flip bits on
		} else {
			val &= ^byte(0x04) // flip bits off
		}

		if cmd.AlarmEvent4.AlarmTemperatureExceeded4 {
			val |= byte(0x08) // flip bits on
		} else {
			val &= ^byte(0x08) // flip bits off
		}

		if cmd.AlarmEvent4.AlarmStillActive4 {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	{
		var val byte

		if cmd.AlarmEvent5.ExternalAlarm5 {
			val |= byte(0x02) // flip bits on
		} else {
			val &= ^byte(0x02) // flip bits off
		}

		if cmd.AlarmEvent5.SensorError5 {
			val |= byte(0x04) // flip bits on
		} else {
			val &= ^byte(0x04) // flip bits off
		}

		if cmd.AlarmEvent5.AlarmTemperatureExceeded5 {
			val |= byte(0x08) // flip bits on
		} else {
			val &= ^byte(0x08) // flip bits off
		}

		if cmd.AlarmEvent5.AlarmStillActive5 {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	return
}