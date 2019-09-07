// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package chimneyfan

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandStopTimeSet cc.CommandID = 0x13

func init() {
	gob.Register(StopTimeSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x2A),
		Command:      cc.CommandID(0x13),
		Version:      1,
	}, NewStopTimeSet)
}

func NewStopTimeSet() cc.Command {
	return &StopTimeSet{}
}

// <no value>
type StopTimeSet struct {
	Time byte
}

func (cmd StopTimeSet) CommandClassID() cc.CommandClassID {
	return 0x2A
}

func (cmd StopTimeSet) CommandID() cc.CommandID {
	return CommandStopTimeSet
}

func (cmd StopTimeSet) CommandIDString() string {
	return "CHIMNEY_FAN_STOP_TIME_SET"
}

func (cmd *StopTimeSet) UnmarshalBinary(data []byte) error {
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

	cmd.Time = payload[i]
	i++

	return nil
}

func (cmd *StopTimeSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.Time)

	return
}
