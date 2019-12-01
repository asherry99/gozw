// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package metertblmonitor

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandMeterTblReport cc.CommandID = 0x06

func init() {
	gob.Register(MeterTblReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x3D),
		Command:      cc.CommandID(0x06),
		Version:      1,
	}, NewMeterTblReport)
}

func NewMeterTblReport() cc.Command {
	return &MeterTblReport{}
}

// <no value>
type MeterTblReport struct {
	Properties1 struct {
		MeterType byte

		RateType byte
	}

	Properties2 struct {
		PayMeter byte
	}

	DatasetSupported uint32

	DatasetHistorySupported uint32

	DataHistorySupported uint32
}

func (cmd MeterTblReport) CommandClassID() cc.CommandClassID {
	return 0x3D
}

func (cmd MeterTblReport) CommandID() cc.CommandID {
	return CommandMeterTblReport
}

func (cmd MeterTblReport) CommandIDString() string {
	return "METER_TBL_REPORT"
}

func (cmd *MeterTblReport) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.MeterType = (payload[i] & 0x3F)

	cmd.Properties1.RateType = (payload[i] & 0xC0) >> 6

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties2.PayMeter = (payload[i] & 0x0F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DatasetSupported = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DatasetHistorySupported = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DataHistorySupported = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	return nil
}

func (cmd *MeterTblReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.MeterType) & byte(0x3F)

		val |= (cmd.Properties1.RateType << byte(6)) & byte(0xC0)

		payload = append(payload, val)
	}

	{
		var val byte

		val |= (cmd.Properties2.PayMeter) & byte(0x0F)

		payload = append(payload, val)
	}

	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.DatasetSupported)
		if buf[0] != 0 {
			return nil, errors.New("BIT_24 value overflow")
		}
		payload = append(payload, buf[1:4]...)
	}

	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.DatasetHistorySupported)
		if buf[0] != 0 {
			return nil, errors.New("BIT_24 value overflow")
		}
		payload = append(payload, buf[1:4]...)
	}

	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.DataHistorySupported)
		if buf[0] != 0 {
			return nil, errors.New("BIT_24 value overflow")
		}
		payload = append(payload, buf[1:4]...)
	}

	return
}