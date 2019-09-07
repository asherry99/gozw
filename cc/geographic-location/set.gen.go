// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package geographiclocation

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandSet cc.CommandID = 0x01

func init() {
	gob.Register(Set{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x8C),
		Command:      cc.CommandID(0x01),
		Version:      1,
	}, NewSet)
}

func NewSet() cc.Command {
	return &Set{}
}

// <no value>
type Set struct {
	LongitudeDegrees byte

	Level struct {
		LongitudeMinutes byte

		LongSign bool
	}

	LatitudeDegrees byte

	Level2 struct {
		LatitudeMinutes byte

		LatSign bool
	}
}

func (cmd Set) CommandClassID() cc.CommandClassID {
	return 0x8C
}

func (cmd Set) CommandID() cc.CommandID {
	return CommandSet
}

func (cmd Set) CommandIDString() string {
	return "GEOGRAPHIC_LOCATION_SET"
}

func (cmd *Set) UnmarshalBinary(data []byte) error {
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

	cmd.LongitudeDegrees = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Level.LongitudeMinutes = (payload[i] & 0x7F)

	cmd.Level.LongSign = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.LatitudeDegrees = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Level2.LatitudeMinutes = (payload[i] & 0x7F)

	cmd.Level2.LatSign = payload[i]&0x80 == 0x80

	i += 1

	return nil
}

func (cmd *Set) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.LongitudeDegrees)

	{
		var val byte

		val |= (cmd.Level.LongitudeMinutes) & byte(0x7F)

		if cmd.Level.LongSign {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.LatitudeDegrees)

	{
		var val byte

		val |= (cmd.Level2.LatitudeMinutes) & byte(0x7F)

		if cmd.Level2.LatSign {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	return
}
