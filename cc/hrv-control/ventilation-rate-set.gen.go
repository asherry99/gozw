// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package hrvcontrol

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandVentilationRateSet cc.CommandID = 0x07

func init() {
	gob.Register(VentilationRateSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x39),
		Command:      cc.CommandID(0x07),
		Version:      1,
	}, NewVentilationRateSet)
}

func NewVentilationRateSet() cc.Command {
	return &VentilationRateSet{}
}

// <no value>
type VentilationRateSet struct {
	VentilationRate byte
}

func (cmd VentilationRateSet) CommandClassID() cc.CommandClassID {
	return 0x39
}

func (cmd VentilationRateSet) CommandID() cc.CommandID {
	return CommandVentilationRateSet
}

func (cmd VentilationRateSet) CommandIDString() string {
	return "HRV_CONTROL_VENTILATION_RATE_SET"
}

func (cmd *VentilationRateSet) UnmarshalBinary(data []byte) error {
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

	cmd.VentilationRate = payload[i]
	i++

	return nil
}

func (cmd *VentilationRateSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.VentilationRate)

	return
}