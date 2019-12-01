// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package inclusioncontroller

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandComplete cc.CommandID = 0x02

func init() {
	gob.Register(Complete{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x74),
		Command:      cc.CommandID(0x02),
		Version:      1,
	}, NewComplete)
}

func NewComplete() cc.Command {
	return &Complete{}
}

// <no value>
type Complete struct {
	StepId byte

	Status byte
}

func (cmd Complete) CommandClassID() cc.CommandClassID {
	return 0x74
}

func (cmd Complete) CommandID() cc.CommandID {
	return CommandComplete
}

func (cmd Complete) CommandIDString() string {
	return "COMPLETE"
}

func (cmd *Complete) UnmarshalBinary(data []byte) error {
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

	cmd.StepId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Status = payload[i]
	i++

	return nil
}

func (cmd *Complete) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.StepId)

	payload = append(payload, cmd.Status)

	return
}