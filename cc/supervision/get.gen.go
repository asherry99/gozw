// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package supervision

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandGet cc.CommandID = 0x01

func init() {
	gob.Register(Get{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x6C),
		Command:      cc.CommandID(0x01),
		Version:      1,
	}, NewGet)
}

func NewGet() cc.Command {
	return &Get{}
}

// <no value>
type Get struct {
	Properties1 struct {
		SessionId byte

		StatusUpdates bool
	}

	EncapsulatedCommandLength byte

	EncapsulatedCommand []byte
}

func (cmd Get) CommandClassID() cc.CommandClassID {
	return 0x6C
}

func (cmd Get) CommandID() cc.CommandID {
	return CommandGet
}

func (cmd Get) CommandIDString() string {
	return "SUPERVISION_GET"
}

func (cmd *Get) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.SessionId = (payload[i] & 0x3F)

	cmd.Properties1.StatusUpdates = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.EncapsulatedCommandLength = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	{
		length := (payload[1+2]) & 0xFF
		cmd.EncapsulatedCommand = payload[i : i+int(length)]
		i += int(length)
	}

	return nil
}

func (cmd *Get) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.SessionId) & byte(0x3F)

		if cmd.Properties1.StatusUpdates {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.EncapsulatedCommandLength)

	if cmd.EncapsulatedCommand != nil && len(cmd.EncapsulatedCommand) > 0 {
		payload = append(payload, cmd.EncapsulatedCommand...)
	}

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}