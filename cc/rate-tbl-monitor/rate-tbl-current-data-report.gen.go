// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package ratetblmonitor

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandRateTblCurrentDataReport cc.CommandID = 0x08

func init() {
	gob.Register(RateTblCurrentDataReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x49),
		Command:      cc.CommandID(0x08),
		Version:      1,
	}, NewRateTblCurrentDataReport)
}

func NewRateTblCurrentDataReport() cc.Command {
	return &RateTblCurrentDataReport{}
}

// <no value>
type RateTblCurrentDataReport struct {
	ReportsToFollow byte

	RateParameterSetId byte

	Dataset uint32

	Year uint16

	Month byte

	Day byte

	HourLocalTime byte

	MinuteLocalTime byte

	SecondLocalTime byte

	Vg []RateTblCurrentDataReportVg
}

type RateTblCurrentDataReportVg struct {
	Properties1 struct {
		CurrentScale byte

		CurrentPrecision byte
	}

	CurrentValue uint32
}

func (cmd RateTblCurrentDataReport) CommandClassID() cc.CommandClassID {
	return 0x49
}

func (cmd RateTblCurrentDataReport) CommandID() cc.CommandID {
	return CommandRateTblCurrentDataReport
}

func (cmd RateTblCurrentDataReport) CommandIDString() string {
	return "RATE_TBL_CURRENT_DATA_REPORT"
}

func (cmd *RateTblCurrentDataReport) UnmarshalBinary(data []byte) error {
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

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.RateParameterSetId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Dataset = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Year = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Month = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Day = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.HourLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MinuteLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SecondLocalTime = payload[i]
	i++

	for i < len(payload) {

		vg := RateTblCurrentDataReportVg{}

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		vg.Properties1.CurrentScale = (payload[i] & 0x1F)

		vg.Properties1.CurrentPrecision = (payload[i] & 0xE0) >> 5

		i += 1

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		currentValue := binary.BigEndian.Uint32(payload[i : i+4])
		i += 4

		// struct byte fields are assigned to the variant group when computed

		vg.CurrentValue = currentValue

		cmd.Vg = append(cmd.Vg, vg)
	}

	return nil
}

func (cmd *RateTblCurrentDataReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.ReportsToFollow)

	payload = append(payload, cmd.RateParameterSetId)

	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.Dataset)
		if buf[0] != 0 {
			return nil, errors.New("BIT_24 value overflow")
		}
		payload = append(payload, buf[1:4]...)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.Year)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.Month)

	payload = append(payload, cmd.Day)

	payload = append(payload, cmd.HourLocalTime)

	payload = append(payload, cmd.MinuteLocalTime)

	payload = append(payload, cmd.SecondLocalTime)

	for _, vg := range cmd.Vg {

		{
			var val byte

			val |= (vg.Properties1.CurrentScale) & byte(0x1F)

			val |= (vg.Properties1.CurrentPrecision << byte(5)) & byte(0xE0)

			payload = append(payload, val)
		}

		{
			buf := make([]byte, 4)
			binary.BigEndian.PutUint32(buf, vg.CurrentValue)
			payload = append(payload, buf...)
		}

	}

	return
}