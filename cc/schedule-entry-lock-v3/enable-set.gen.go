// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package scheduleentrylockv3

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandEnableSet cc.CommandID = 0x01

func init() {
	gob.Register(EnableSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x4E),
		Command:      cc.CommandID(0x01),
		Version:      3,
	}, NewEnableSet)
}

func NewEnableSet() cc.Command {
	return &EnableSet{}
}

// <no value>
type EnableSet struct {
	UserIdentifier byte

	Enabled byte
}

func (cmd EnableSet) CommandClassID() cc.CommandClassID {
	return 0x4E
}

func (cmd EnableSet) CommandID() cc.CommandID {
	return CommandEnableSet
}

func (cmd EnableSet) CommandIDString() string {
	return "SCHEDULE_ENTRY_LOCK_ENABLE_SET"
}

func (cmd *EnableSet) UnmarshalBinary(data []byte) error {
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

	cmd.UserIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Enabled = payload[i]
	i++

	return nil
}

func (cmd *EnableSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.UserIdentifier)

	payload = append(payload, cmd.Enabled)

	return
}
