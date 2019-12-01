// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package zipportal

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandGatewayConfigurationStatus cc.CommandID = 0x02

func init() {
	gob.Register(GatewayConfigurationStatus{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x61),
		Command:      cc.CommandID(0x02),
		Version:      1,
	}, NewGatewayConfigurationStatus)
}

func NewGatewayConfigurationStatus() cc.Command {
	return &GatewayConfigurationStatus{}
}

// <no value>
type GatewayConfigurationStatus struct {
	Status byte
}

func (cmd GatewayConfigurationStatus) CommandClassID() cc.CommandClassID {
	return 0x61
}

func (cmd GatewayConfigurationStatus) CommandID() cc.CommandID {
	return CommandGatewayConfigurationStatus
}

func (cmd GatewayConfigurationStatus) CommandIDString() string {
	return "GATEWAY_CONFIGURATION_STATUS"
}

func (cmd *GatewayConfigurationStatus) UnmarshalBinary(data []byte) error {
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

	cmd.Status = payload[i]
	i++

	return nil
}

func (cmd *GatewayConfigurationStatus) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.Status)

	return
}