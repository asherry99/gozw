// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package configurationv3

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandPropertiesReport cc.CommandID = 0x0F

func init() {
	gob.Register(PropertiesReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x70),
		Command:      cc.CommandID(0x0F),
		Version:      3,
	}, NewPropertiesReport)
}

func NewPropertiesReport() cc.Command {
	return &PropertiesReport{}
}

// <no value>
type PropertiesReport struct {
	ParameterNumber uint16

	Properties1 struct {
		Size byte

		Format byte
	}

	MinValue []byte

	MaxValue []byte

	DefaultValue []byte

	NextParameterNumber uint16
}

func (cmd PropertiesReport) CommandClassID() cc.CommandClassID {
	return 0x70
}

func (cmd PropertiesReport) CommandID() cc.CommandID {
	return CommandPropertiesReport
}

func (cmd PropertiesReport) CommandIDString() string {
	return "CONFIGURATION_PROPERTIES_REPORT"
}

func (cmd *PropertiesReport) UnmarshalBinary(data []byte) error {
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

	cmd.ParameterNumber = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.Size = (payload[i] & 0x07)

	cmd.Properties1.Format = (payload[i] & 0x38) >> 3

	i += 1

	if len(payload) <= i {
		return nil // field is optional
	}

	{
		length := (payload[1+2]) & 0x07
		cmd.MinValue = payload[i : i+int(length)]
		i += int(length)
	}

	if len(payload) <= i {
		return nil // field is optional
	}

	{
		length := (payload[1+2]) & 0x07
		cmd.MaxValue = payload[i : i+int(length)]
		i += int(length)
	}

	if len(payload) <= i {
		return nil // field is optional
	}

	{
		length := (payload[1+2]) & 0x07
		cmd.DefaultValue = payload[i : i+int(length)]
		i += int(length)
	}

	cmd.NextParameterNumber = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}

func (cmd *PropertiesReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ParameterNumber)
		payload = append(payload, buf...)
	}

	{
		var val byte

		val |= (cmd.Properties1.Size) & byte(0x07)

		val |= (cmd.Properties1.Format << byte(3)) & byte(0x38)

		payload = append(payload, val)
	}

	if cmd.MinValue != nil && len(cmd.MinValue) > 0 {
		payload = append(payload, cmd.MinValue...)
	}

	if cmd.MaxValue != nil && len(cmd.MaxValue) > 0 {
		payload = append(payload, cmd.MaxValue...)
	}

	if cmd.DefaultValue != nil && len(cmd.DefaultValue) > 0 {
		payload = append(payload, cmd.DefaultValue...)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.NextParameterNumber)
		payload = append(payload, buf...)
	}

	return
}
