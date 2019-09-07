// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package humiditycontrolsetpointv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandScaleSupportedGet cc.CommandID = 0x06

func init() {
	gob.Register(ScaleSupportedGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x64),
		Command:      cc.CommandID(0x06),
		Version:      2,
	}, NewScaleSupportedGet)
}

func NewScaleSupportedGet() cc.Command {
	return &ScaleSupportedGet{}
}

// <no value>
type ScaleSupportedGet struct {
	Properties1 struct {
		SetpointType byte
	}
}

func (cmd ScaleSupportedGet) CommandClassID() cc.CommandClassID {
	return 0x64
}

func (cmd ScaleSupportedGet) CommandID() cc.CommandID {
	return CommandScaleSupportedGet
}

func (cmd ScaleSupportedGet) CommandIDString() string {
	return "HUMIDITY_CONTROL_SETPOINT_SCALE_SUPPORTED_GET"
}

func (cmd *ScaleSupportedGet) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.SetpointType = (payload[i] & 0x0F)

	i += 1

	return nil
}

func (cmd *ScaleSupportedGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.SetpointType) & byte(0x0F)

		payload = append(payload, val)
	}

	return
}
