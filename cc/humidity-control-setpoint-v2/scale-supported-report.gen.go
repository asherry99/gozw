// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package humiditycontrolsetpointv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandScaleSupportedReport cc.CommandID = 0x07

func init() {
	gob.Register(ScaleSupportedReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x64),
		Command:      cc.CommandID(0x07),
		Version:      2,
	}, NewScaleSupportedReport)
}

func NewScaleSupportedReport() cc.Command {
	return &ScaleSupportedReport{}
}

// <no value>
type ScaleSupportedReport struct {
	Properties1 struct {
		ScaleBitMask byte
	}
}

func (cmd ScaleSupportedReport) CommandClassID() cc.CommandClassID {
	return 0x64
}

func (cmd ScaleSupportedReport) CommandID() cc.CommandID {
	return CommandScaleSupportedReport
}

func (cmd ScaleSupportedReport) CommandIDString() string {
	return "HUMIDITY_CONTROL_SETPOINT_SCALE_SUPPORTED_REPORT"
}

func (cmd *ScaleSupportedReport) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.ScaleBitMask = (payload[i] & 0x0F)

	i += 1

	return nil
}

func (cmd *ScaleSupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.ScaleBitMask) & byte(0x0F)

		payload = append(payload, val)
	}

	return
}