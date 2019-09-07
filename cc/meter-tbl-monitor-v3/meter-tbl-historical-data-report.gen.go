// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package metertblmonitorv3

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandMeterTblHistoricalDataReport cc.CommandID = 0x0F

func init() {
	gob.Register(MeterTblHistoricalDataReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x3D),
		Command:      cc.CommandID(0x0F),
		Version:      3,
	}, NewMeterTblHistoricalDataReport)
}

func NewMeterTblHistoricalDataReport() cc.Command {
	return &MeterTblHistoricalDataReport{}
}

// <no value>
type MeterTblHistoricalDataReport struct {
	ReportsToFollow byte

	Properties1 struct {
		RateType byte

		OperatingStatusIndication bool
	}

	Dataset uint32

	Year uint16

	Month byte

	Day byte

	HourLocalTime byte

	MinuteLocalTime byte

	SecondLocalTime byte

	Vg []MeterTblHistoricalDataReportVg
}

type MeterTblHistoricalDataReportVg struct {
	Properties1 struct {
		HistoricalScale byte

		HistoricalPrecision byte
	}

	HistoricalValue uint32
}

func (cmd MeterTblHistoricalDataReport) CommandClassID() cc.CommandClassID {
	return 0x3D
}

func (cmd MeterTblHistoricalDataReport) CommandID() cc.CommandID {
	return CommandMeterTblHistoricalDataReport
}

func (cmd MeterTblHistoricalDataReport) CommandIDString() string {
	return "METER_TBL_HISTORICAL_DATA_REPORT"
}

func (cmd *MeterTblHistoricalDataReport) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.RateType = (payload[i] & 0x03)

	cmd.Properties1.OperatingStatusIndication = payload[i]&0x80 == 0x80

	i += 1

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

		vg := MeterTblHistoricalDataReportVg{}

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		vg.Properties1.HistoricalScale = (payload[i] & 0x1F)

		vg.Properties1.HistoricalPrecision = (payload[i] & 0xE0) >> 5

		i += 1

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		historicalValue := binary.BigEndian.Uint32(payload[i : i+4])
		i += 4

		// struct byte fields are assigned to the variant group when computed

		vg.HistoricalValue = historicalValue

		cmd.Vg = append(cmd.Vg, vg)
	}

	return nil
}

func (cmd *MeterTblHistoricalDataReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.ReportsToFollow)

	{
		var val byte

		val |= (cmd.Properties1.RateType) & byte(0x03)

		if cmd.Properties1.OperatingStatusIndication {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

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

			val |= (vg.Properties1.HistoricalScale) & byte(0x1F)

			val |= (vg.Properties1.HistoricalPrecision << byte(5)) & byte(0xE0)

			payload = append(payload, val)
		}

		{
			buf := make([]byte, 4)
			binary.BigEndian.PutUint32(buf, vg.HistoricalValue)
			payload = append(payload, buf...)
		}

	}

	return
}
