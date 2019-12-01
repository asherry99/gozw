// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package barrieroperator

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandSignalGet cc.CommandID = 0x07

func init() {
	gob.Register(SignalGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x66),
		Command:      cc.CommandID(0x07),
		Version:      1,
	}, NewSignalGet)
}

func NewSignalGet() cc.Command {
	return &SignalGet{}
}

// <no value>
type SignalGet struct {
	SubsystemType byte
}

func (cmd SignalGet) CommandClassID() cc.CommandClassID {
	return 0x66
}

func (cmd SignalGet) CommandID() cc.CommandID {
	return CommandSignalGet
}

func (cmd SignalGet) CommandIDString() string {
	return "BARRIER_OPERATOR_SIGNAL_GET"
}

func (cmd *SignalGet) UnmarshalBinary(data []byte) error {
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

	cmd.SubsystemType = payload[i]
	i++

	return nil
}

func (cmd *SignalGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SubsystemType)

	return
}