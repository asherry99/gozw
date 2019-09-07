// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package windowcovering

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandStopLevelChange cc.CommandID = 0x07

func init() {
	gob.Register(StopLevelChange{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x6A),
		Command:      cc.CommandID(0x07),
		Version:      1,
	}, NewStopLevelChange)
}

func NewStopLevelChange() cc.Command {
	return &StopLevelChange{}
}

// <no value>
type StopLevelChange struct {
	ParameterId byte
}

func (cmd StopLevelChange) CommandClassID() cc.CommandClassID {
	return 0x6A
}

func (cmd StopLevelChange) CommandID() cc.CommandID {
	return CommandStopLevelChange
}

func (cmd StopLevelChange) CommandIDString() string {
	return "WINDOW_COVERING_STOP_LEVEL_CHANGE"
}

func (cmd *StopLevelChange) UnmarshalBinary(data []byte) error {
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

	cmd.ParameterId = payload[i]
	i++

	return nil
}

func (cmd *StopLevelChange) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.ParameterId)

	return
}
