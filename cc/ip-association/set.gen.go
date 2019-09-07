// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package ipassociation

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandSet cc.CommandID = 0x01

func init() {
	gob.Register(Set{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x5C),
		Command:      cc.CommandID(0x01),
		Version:      1,
	}, NewSet)
}

func NewSet() cc.Command {
	return &Set{}
}

// <no value>
type Set struct {
	GroupingIdentifier byte

	Ipv6Address []byte

	EndPoint byte
}

func (cmd Set) CommandClassID() cc.CommandClassID {
	return 0x5C
}

func (cmd Set) CommandID() cc.CommandID {
	return CommandSet
}

func (cmd Set) CommandIDString() string {
	return "IP_ASSOCIATION_SET"
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

	cmd.GroupingIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Ipv6Address = payload[i : i+16]

	i += 16

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.EndPoint = payload[i]
	i++

	return nil
}

func (cmd *Set) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.GroupingIdentifier)

	if paramLen := len(cmd.Ipv6Address); paramLen > 16 {
		return nil, errors.New("Length overflow in array parameter Ipv6Address")
	}

	payload = append(payload, cmd.Ipv6Address...)

	payload = append(payload, cmd.EndPoint)

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}
