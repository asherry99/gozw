// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package thermostatheating

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandSetpointReport cc.CommandID = 0x06

func init() {
	gob.Register(SetpointReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x38),
		Command:      cc.CommandID(0x06),
		Version:      1,
	}, NewSetpointReport)
}

func NewSetpointReport() cc.Command {
	return &SetpointReport{}
}

// <no value>
type SetpointReport struct {
	SetpointNr byte

	Properties1 struct {
		Size byte

		Scale byte

		Precision byte
	}

	Value []byte
}

func (cmd SetpointReport) CommandClassID() cc.CommandClassID {
	return 0x38
}

func (cmd SetpointReport) CommandID() cc.CommandID {
	return CommandSetpointReport
}

func (cmd SetpointReport) CommandIDString() string {
	return "THERMOSTAT_HEATING_SETPOINT_REPORT"
}

func (cmd *SetpointReport) UnmarshalBinary(data []byte) error {
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

	cmd.SetpointNr = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.Size = (payload[i] & 0x07)

	cmd.Properties1.Scale = (payload[i] & 0x18) >> 3

	cmd.Properties1.Precision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return nil // field is optional
	}

	{
		length := (payload[1+2]) & 0x07
		cmd.Value = payload[i : i+int(length)]
		i += int(length)
	}

	return nil
}

func (cmd *SetpointReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SetpointNr)

	{
		var val byte

		val |= (cmd.Properties1.Size) & byte(0x07)

		val |= (cmd.Properties1.Scale << byte(3)) & byte(0x18)

		val |= (cmd.Properties1.Precision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	if cmd.Value != nil && len(cmd.Value) > 0 {
		payload = append(payload, cmd.Value...)
	}

	return
}
