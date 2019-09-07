// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package multicmd

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandEncap cc.CommandID = 0x01

func init() {
	gob.Register(Encap{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x8F),
		Command:      cc.CommandID(0x01),
		Version:      1,
	}, NewEncap)
}

func NewEncap() cc.Command {
	return &Encap{}
}

// <no value>
type Encap struct {
	NumberOfCommands byte

	EncapsulatedCommand []EncapEncapsulatedCommand
}

type EncapEncapsulatedCommand struct {
	CommandLength byte

	CommandClass byte

	Command byte

	Data []byte
}

func (cmd Encap) CommandClassID() cc.CommandClassID {
	return 0x8F
}

func (cmd Encap) CommandID() cc.CommandID {
	return CommandEncap
}

func (cmd Encap) CommandIDString() string {
	return "MULTI_CMD_ENCAP"
}

func (cmd *Encap) UnmarshalBinary(data []byte) error {
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

	cmd.NumberOfCommands = payload[i]
	i++

	for i < len(payload) {

		encapsulatedCommand := EncapEncapsulatedCommand{}

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		commandLength := payload[i]
		i++

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		commandClass := payload[i]
		i++

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		command := payload[i]
		i++

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		length := (payload[0+2]) & 0xFF
		data := payload[i : i+int(length)]
		i += int(length)

		encapsulatedCommand.CommandLength = commandLength

		encapsulatedCommand.CommandClass = commandClass

		encapsulatedCommand.Command = command

		encapsulatedCommand.Data = data

		cmd.EncapsulatedCommand = append(cmd.EncapsulatedCommand, encapsulatedCommand)
	}

	return nil
}

func (cmd *Encap) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.NumberOfCommands)

	for _, vg := range cmd.EncapsulatedCommand {

		payload = append(payload, vg.CommandLength)

		payload = append(payload, vg.CommandClass)

		payload = append(payload, vg.Command)

		if vg.Data != nil && len(vg.Data) > 0 {
			payload = append(payload, vg.Data...)
		}

	}

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}
