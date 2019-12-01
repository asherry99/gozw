// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package associationcommandconfiguration

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandRecordsSupportedReport cc.CommandID = 0x02

func init() {
	gob.Register(RecordsSupportedReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x9B),
		Command:      cc.CommandID(0x02),
		Version:      1,
	}, NewRecordsSupportedReport)
}

func NewRecordsSupportedReport() cc.Command {
	return &RecordsSupportedReport{}
}

// <no value>
type RecordsSupportedReport struct {
	Properties1 struct {
		MaxCommandLength byte

		ConfCmd bool

		Vc bool
	}

	FreeCommandRecords uint16

	MaxCommandRecords uint16
}

func (cmd RecordsSupportedReport) CommandClassID() cc.CommandClassID {
	return 0x9B
}

func (cmd RecordsSupportedReport) CommandID() cc.CommandID {
	return CommandRecordsSupportedReport
}

func (cmd RecordsSupportedReport) CommandIDString() string {
	return "COMMAND_RECORDS_SUPPORTED_REPORT"
}

func (cmd *RecordsSupportedReport) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.MaxCommandLength = (payload[i] & 0xFC) >> 2

	cmd.Properties1.ConfCmd = payload[i]&0x01 == 0x01

	cmd.Properties1.Vc = payload[i]&0x02 == 0x02

	i += 1

	cmd.FreeCommandRecords = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	cmd.MaxCommandRecords = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}

func (cmd *RecordsSupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.MaxCommandLength << byte(2)) & byte(0xFC)

		if cmd.Properties1.ConfCmd {
			val |= byte(0x01) // flip bits on
		} else {
			val &= ^byte(0x01) // flip bits off
		}

		if cmd.Properties1.Vc {
			val |= byte(0x02) // flip bits on
		} else {
			val &= ^byte(0x02) // flip bits off
		}

		payload = append(payload, val)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.FreeCommandRecords)
		payload = append(payload, buf...)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.MaxCommandRecords)
		payload = append(payload, buf...)
	}

	return
}