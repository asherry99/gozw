// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package securitypanelzonesensor

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandStateGet cc.CommandID = 0x05

func init() {
	gob.Register(StateGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x2F),
		Command:      cc.CommandID(0x05),
		Version:      1,
	}, NewStateGet)
}

func NewStateGet() cc.Command {
	return &StateGet{}
}

// <no value>
type StateGet struct {
	ZoneNumber byte

	SensorNumber byte
}

func (cmd StateGet) CommandClassID() cc.CommandClassID {
	return 0x2F
}

func (cmd StateGet) CommandID() cc.CommandID {
	return CommandStateGet
}

func (cmd StateGet) CommandIDString() string {
	return "SECURITY_PANEL_ZONE_SENSOR_STATE_GET"
}

func (cmd *StateGet) UnmarshalBinary(data []byte) error {
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

	cmd.ZoneNumber = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SensorNumber = payload[i]
	i++

	return nil
}

func (cmd *StateGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.ZoneNumber)

	payload = append(payload, cmd.SensorNumber)

	return
}