// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package centralscenev3

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandConfigurationSet cc.CommandID = 0x04

func init() {
	gob.Register(ConfigurationSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x5B),
		Command:      cc.CommandID(0x04),
		Version:      3,
	}, NewConfigurationSet)
}

func NewConfigurationSet() cc.Command {
	return &ConfigurationSet{}
}

// <no value>
type ConfigurationSet struct {
	Properties1 struct {
		SlowRefresh bool
	}
}

func (cmd ConfigurationSet) CommandClassID() cc.CommandClassID {
	return 0x5B
}

func (cmd ConfigurationSet) CommandID() cc.CommandID {
	return CommandConfigurationSet
}

func (cmd ConfigurationSet) CommandIDString() string {
	return "CENTRAL_SCENE_CONFIGURATION_SET"
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

	cmd.Properties1.SlowRefresh = payload[i]&0x80 == 0x80

	i += 1

	return nil
}

func (cmd *ConfigurationSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		if cmd.Properties1.SlowRefresh {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	return
}