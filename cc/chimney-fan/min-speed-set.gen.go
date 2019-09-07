// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package chimneyfan

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandMinSpeedSet cc.CommandID = 0x25

func init() {
	gob.Register(MinSpeedSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x2A),
		Command:      cc.CommandID(0x25),
		Version:      1,
	}, NewMinSpeedSet)
}

func NewMinSpeedSet() cc.Command {
	return &MinSpeedSet{}
}

// <no value>
type MinSpeedSet struct {
	MinSpeed byte
}

func (cmd MinSpeedSet) CommandClassID() cc.CommandClassID {
	return 0x2A
}

func (cmd MinSpeedSet) CommandID() cc.CommandID {
	return CommandMinSpeedSet
}

func (cmd MinSpeedSet) CommandIDString() string {
	return "CHIMNEY_FAN_MIN_SPEED_SET"
}

func (cmd *MinSpeedSet) UnmarshalBinary(data []byte) error {
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

	cmd.MinSpeed = payload[i]
	i++

	return nil
}

func (cmd *MinSpeedSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.MinSpeed)

	return
}
