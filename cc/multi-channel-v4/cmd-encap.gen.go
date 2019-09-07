// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package multichannelv4

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandCmdEncap cc.CommandID = 0x0D

func init() {
	gob.Register(CmdEncap{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x60),
		Command:      cc.CommandID(0x0D),
		Version:      4,
	}, NewCmdEncap)
}

func NewCmdEncap() cc.Command {
	return &CmdEncap{}
}

// <no value>
type CmdEncap struct {
	Properties1 struct {
		SourceEndPoint byte
	}

	Properties2 struct {
		DestinationEndPoint byte

		BitAddress bool
	}

	CommandClass byte

	Command byte

	Parameter []byte
}

func (cmd CmdEncap) CommandClassID() cc.CommandClassID {
	return 0x60
}

func (cmd CmdEncap) CommandID() cc.CommandID {
	return CommandCmdEncap
}

func (cmd CmdEncap) CommandIDString() string {
	return "MULTI_CHANNEL_CMD_ENCAP"
}

func (cmd *CmdEncap) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.SourceEndPoint = (payload[i] & 0x7F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties2.DestinationEndPoint = (payload[i] & 0x7F)

	cmd.Properties2.BitAddress = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandClass = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Command = payload[i]
	i++

	if len(payload) <= i {
		return nil
	}

	cmd.Parameter = payload[i:]

	return nil
}

func (cmd *CmdEncap) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.SourceEndPoint) & byte(0x7F)

		payload = append(payload, val)
	}

	{
		var val byte

		val |= (cmd.Properties2.DestinationEndPoint) & byte(0x7F)

		if cmd.Properties2.BitAddress {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.CommandClass)

	payload = append(payload, cmd.Command)

	payload = append(payload, cmd.Parameter...)

	return
}
