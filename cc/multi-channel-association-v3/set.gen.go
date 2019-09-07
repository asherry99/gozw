// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package multichannelassociationv3

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandSet cc.CommandID = 0x01

func init() {
	gob.Register(Set{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x8E),
		Command:      cc.CommandID(0x01),
		Version:      3,
	}, NewSet)
}

func NewSet() cc.Command {
	return &Set{}
}

// <no value>
type Set struct {
	GroupingIdentifier byte

	NodeId []byte

	Vg []SetVg
}

type SetVg struct {
	MultiChannelNodeId byte

	Properties1 struct {
		EndPoint byte

		BitAddress bool
	}
}

func (cmd Set) CommandClassID() cc.CommandClassID {
	return 0x8E
}

func (cmd Set) CommandID() cc.CommandID {
	return CommandSet
}

func (cmd Set) CommandIDString() string {
	return "MULTI_CHANNEL_ASSOCIATION_SET"
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

	{
		fieldStart := i
		for ; i < len(payload) && payload[i] != 0x00; i++ {
		}
		cmd.NodeId = payload[fieldStart:i]
	}

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	i += 1 // skipping MARKER
	if len(payload) <= i {
		return nil
	}

	for i < len(payload) {

		vg := SetVg{}

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		multiChannelNodeId := payload[i]
		i++

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		vg.Properties1.EndPoint = (payload[i] & 0x7F)

		vg.Properties1.BitAddress = payload[i]&0x80 == 0x80

		i += 1

		vg.MultiChannelNodeId = multiChannelNodeId

		// struct byte fields are assigned to the variant group when computed

		cmd.Vg = append(cmd.Vg, vg)
	}

	return nil
}

func (cmd *Set) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.GroupingIdentifier)

	{
		if cmd.NodeId != nil && len(cmd.NodeId) > 0 {
			payload = append(payload, cmd.NodeId...)
		}
		payload = append(payload, 0x00)
	}

	payload = append(payload, 0x00) // marker

	for _, vg := range cmd.Vg {

		payload = append(payload, vg.MultiChannelNodeId)

		{
			var val byte

			val |= (vg.Properties1.EndPoint) & byte(0x7F)

			if vg.Properties1.BitAddress {
				val |= byte(0x80) // flip bits on
			} else {
				val &= ^byte(0x80) // flip bits off
			}

			payload = append(payload, val)
		}

	}

	return
}
