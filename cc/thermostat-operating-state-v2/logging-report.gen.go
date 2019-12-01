// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package thermostatoperatingstatev2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandLoggingReport cc.CommandID = 0x06

func init() {
	gob.Register(LoggingReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x42),
		Command:      cc.CommandID(0x06),
		Version:      2,
	}, NewLoggingReport)
}

func NewLoggingReport() cc.Command {
	return &LoggingReport{}
}

// <no value>
type LoggingReport struct {
	ReportsToFollow byte

	Vg1 []LoggingReportVg1
}

type LoggingReportVg1 struct {
	Properties1 struct {
		OperatingStateLogType byte
	}

	UsageTodayhours byte

	UsageTodayminutes byte

	UsageYesterdayhours byte

	UsageYesterdayminutes byte
}

func (cmd LoggingReport) CommandClassID() cc.CommandClassID {
	return 0x42
}

func (cmd LoggingReport) CommandID() cc.CommandID {
	return CommandLoggingReport
}

func (cmd LoggingReport) CommandIDString() string {
	return "THERMOSTAT_OPERATING_STATE_LOGGING_REPORT"
}

func (cmd *LoggingReport) UnmarshalBinary(data []byte) error {
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

	cmd.ReportsToFollow = payload[i]
	i++

	for i < len(payload) {

		vg1 := LoggingReportVg1{}

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		vg1.Properties1.OperatingStateLogType = (payload[i] & 0x0F)

		i += 1

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		usageTodayhours := payload[i]
		i++

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		usageTodayminutes := payload[i]
		i++

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		usageYesterdayhours := payload[i]
		i++

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		usageYesterdayminutes := payload[i]
		i++

		// struct byte fields are assigned to the variant group when computed

		vg1.UsageTodayhours = usageTodayhours

		vg1.UsageTodayminutes = usageTodayminutes

		vg1.UsageYesterdayhours = usageYesterdayhours

		vg1.UsageYesterdayminutes = usageYesterdayminutes

		cmd.Vg1 = append(cmd.Vg1, vg1)
	}

	return nil
}

func (cmd *LoggingReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.ReportsToFollow)

	for _, vg := range cmd.Vg1 {

		{
			var val byte

			val |= (vg.Properties1.OperatingStateLogType) & byte(0x0F)

			payload = append(payload, val)
		}

		payload = append(payload, vg.UsageTodayhours)

		payload = append(payload, vg.UsageTodayminutes)

		payload = append(payload, vg.UsageYesterdayhours)

		payload = append(payload, vg.UsageYesterdayminutes)

	}

	return
}