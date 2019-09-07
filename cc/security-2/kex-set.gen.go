// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package security2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandKexSet cc.CommandID = 0x06

func init() {
	gob.Register(KexSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x9F),
		Command:      cc.CommandID(0x06),
		Version:      1,
	}, NewKexSet)
}

func NewKexSet() cc.Command {
	return &KexSet{}
}

// <no value>
type KexSet struct {
	Properties1 struct {
		Echo bool

		RequestCsa bool
	}

	SelectedKexScheme byte

	SelectedEcdhProfile byte

	GrantedKeys []byte
}

func (cmd KexSet) CommandClassID() cc.CommandClassID {
	return 0x9F
}

func (cmd KexSet) CommandID() cc.CommandID {
	return CommandKexSet
}

func (cmd KexSet) CommandIDString() string {
	return "KEX_SET"
}

func (cmd *KexSet) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.Echo = payload[i]&0x01 == 0x01

	cmd.Properties1.RequestCsa = payload[i]&0x02 == 0x02

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SelectedKexScheme = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SelectedEcdhProfile = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GrantedKeys = payload[i:]

	return nil
}

func (cmd *KexSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		if cmd.Properties1.Echo {
			val |= byte(0x01) // flip bits on
		} else {
			val &= ^byte(0x01) // flip bits off
		}

		if cmd.Properties1.RequestCsa {
			val |= byte(0x02) // flip bits on
		} else {
			val &= ^byte(0x02) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.SelectedKexScheme)

	payload = append(payload, cmd.SelectedEcdhProfile)

	payload = append(payload, cmd.GrantedKeys...)

	return
}
