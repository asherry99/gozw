// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package hrvstatus

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandGet cc.CommandID = 0x01

func init() {
	gob.Register(Get{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x37),
		Command:      cc.CommandID(0x01),
		Version:      1,
	}, NewGet)
}

func NewGet() cc.Command {
	return &Get{}
}

// <no value>
type Get struct {
	StatusParameter byte
}

func (cmd Get) CommandClassID() cc.CommandClassID {
	return 0x37
}

func (cmd Get) CommandID() cc.CommandID {
	return CommandGet
}

func (cmd Get) CommandIDString() string {
	return "HRV_STATUS_GET"
}

func (cmd *Get) UnmarshalBinary(data []byte) error {
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

	cmd.StatusParameter = payload[i]
	i++

	return nil
}

func (cmd *Get) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.StatusParameter)

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}
