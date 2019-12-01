// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package metertblmonitorv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandMeterTblTablePointAdmNoReport cc.CommandID = 0x02

func init() {
	gob.Register(MeterTblTablePointAdmNoReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x3D),
		Command:      cc.CommandID(0x02),
		Version:      2,
	}, NewMeterTblTablePointAdmNoReport)
}

func NewMeterTblTablePointAdmNoReport() cc.Command {
	return &MeterTblTablePointAdmNoReport{}
}

// <no value>
type MeterTblTablePointAdmNoReport struct {
	Properties1 struct {
		NumberOfCharacters byte
	}

	MeterPointAdmNumberCharacter []byte
}

func (cmd MeterTblTablePointAdmNoReport) CommandClassID() cc.CommandClassID {
	return 0x3D
}

func (cmd MeterTblTablePointAdmNoReport) CommandID() cc.CommandID {
	return CommandMeterTblTablePointAdmNoReport
}

func (cmd MeterTblTablePointAdmNoReport) CommandIDString() string {
	return "METER_TBL_TABLE_POINT_ADM_NO_REPORT"
}

func (cmd *MeterTblTablePointAdmNoReport) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.NumberOfCharacters = (payload[i] & 0x1F)

	i += 1

	if len(payload) <= i {
		return nil // field is optional
	}

	{
		length := (payload[0+2]) & 0x1F
		cmd.MeterPointAdmNumberCharacter = payload[i : i+int(length)]
		i += int(length)
	}

	return nil
}

func (cmd *MeterTblTablePointAdmNoReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.NumberOfCharacters) & byte(0x1F)

		payload = append(payload, val)
	}

	if cmd.MeterPointAdmNumberCharacter != nil && len(cmd.MeterPointAdmNumberCharacter) > 0 {
		payload = append(payload, cmd.MeterPointAdmNumberCharacter...)
	}

	return
}