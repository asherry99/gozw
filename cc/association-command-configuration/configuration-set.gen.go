// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package associationcommandconfiguration

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandConfigurationSet cc.CommandID = 0x03

func init() {
	gob.Register(ConfigurationSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x9B),
		Command:      cc.CommandID(0x03),
		Version:      1,
	}, NewConfigurationSet)
}

func NewConfigurationSet() cc.Command {
	return &ConfigurationSet{}
}

// <no value>
type ConfigurationSet struct {
	GroupingIdentifier byte

	NodeId byte

	CommandLength byte

	CommandClassIdentifier byte

	CommandIdentifier byte

	CommandByte []byte
}

func (cmd ConfigurationSet) CommandClassID() cc.CommandClassID {
	return 0x9B
}

func (cmd ConfigurationSet) CommandID() cc.CommandID {
	return CommandConfigurationSet
}

func (cmd ConfigurationSet) CommandIDString() string {
	return "COMMAND_CONFIGURATION_SET"
}

func (cmd *ConfigurationSet) UnmarshalBinary(data []byte) error {
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

	cmd.GroupingIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NodeId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandLength = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandClassIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return nil
	}

	cmd.CommandByte = payload[i:]

	return nil
}

func (cmd *ConfigurationSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.GroupingIdentifier)

	payload = append(payload, cmd.NodeId)

	payload = append(payload, cmd.CommandLength)

	payload = append(payload, cmd.CommandClassIdentifier)

	payload = append(payload, cmd.CommandIdentifier)

	payload = append(payload, cmd.CommandByte...)

	return
}