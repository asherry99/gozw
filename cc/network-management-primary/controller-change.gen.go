// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package networkmanagementprimary

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandControllerChange cc.CommandID = 0x01

func init() {
	gob.Register(ControllerChange{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x54),
		Command:      cc.CommandID(0x01),
		Version:      1,
	}, NewControllerChange)
}

func NewControllerChange() cc.Command {
	return &ControllerChange{}
}

// <no value>
type ControllerChange struct {
	SeqNo byte

	Mode byte

	TxOptions []byte
}

func (cmd ControllerChange) CommandClassID() cc.CommandClassID {
	return 0x54
}

func (cmd ControllerChange) CommandID() cc.CommandID {
	return CommandControllerChange
}

func (cmd ControllerChange) CommandIDString() string {
	return "CONTROLLER_CHANGE"
}

func (cmd *ControllerChange) UnmarshalBinary(data []byte) error {
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

	cmd.SeqNo = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	i++ // skipping over reserved bit

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Mode = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.TxOptions = payload[i:]

	return nil
}

func (cmd *ControllerChange) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SeqNo)

	payload = append(payload, cmd.Mode)

	payload = append(payload, cmd.TxOptions...)

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}