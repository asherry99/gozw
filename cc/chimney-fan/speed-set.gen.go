// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package chimneyfan

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandSpeedSet cc.CommandID = 0x04

func init() {
	gob.Register(SpeedSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x2A),
		Command:      cc.CommandID(0x04),
		Version:      1,
	}, NewSpeedSet)
}

func NewSpeedSet() cc.Command {
	return &SpeedSet{}
}

// <no value>
type SpeedSet struct {
	Speed byte
}

func (cmd SpeedSet) CommandClassID() cc.CommandClassID {
	return 0x2A
}

func (cmd SpeedSet) CommandID() cc.CommandID {
	return CommandSpeedSet
}

func (cmd SpeedSet) CommandIDString() string {
	return "CHIMNEY_FAN_SPEED_SET"
}

func (cmd *SpeedSet) UnmarshalBinary(data []byte) error {
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

	cmd.Speed = payload[i]
	i++

	return nil
}

func (cmd *SpeedSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.Speed)

	return
}