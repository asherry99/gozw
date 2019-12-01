// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package centralscenev3

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandNotification cc.CommandID = 0x03

func init() {
	gob.Register(Notification{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x5B),
		Command:      cc.CommandID(0x03),
		Version:      3,
	}, NewNotification)
}

func NewNotification() cc.Command {
	return &Notification{}
}

// <no value>
type Notification struct {
	SequenceNumber byte

	Properties1 struct {
		SlowRefresh bool

		KeyAttributes byte
	}

	SceneNumber byte
}

func (cmd Notification) CommandClassID() cc.CommandClassID {
	return 0x5B
}

func (cmd Notification) CommandID() cc.CommandID {
	return CommandNotification
}

func (cmd Notification) CommandIDString() string {
	return "CENTRAL_SCENE_NOTIFICATION"
}

func (cmd *Notification) UnmarshalBinary(data []byte) error {
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

	cmd.SequenceNumber = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.KeyAttributes = (payload[i] & 0x07)

	cmd.Properties1.SlowRefresh = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SceneNumber = payload[i]
	i++

	return nil
}

func (cmd *Notification) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SequenceNumber)

	{
		var val byte

		val |= (cmd.Properties1.KeyAttributes) & byte(0x07)

		if cmd.Properties1.SlowRefresh {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.SceneNumber)

	return
}