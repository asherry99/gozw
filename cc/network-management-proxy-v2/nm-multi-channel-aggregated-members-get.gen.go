// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package networkmanagementproxyv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandNmMultiChannelAggregatedMembersGet cc.CommandID = 0x09

func init() {
	gob.Register(NmMultiChannelAggregatedMembersGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x52),
		Command:      cc.CommandID(0x09),
		Version:      2,
	}, NewNmMultiChannelAggregatedMembersGet)
}

func NewNmMultiChannelAggregatedMembersGet() cc.Command {
	return &NmMultiChannelAggregatedMembersGet{}
}

// <no value>
type NmMultiChannelAggregatedMembersGet struct {
	SeqNo byte

	Nodeid byte

	Properties1 struct {
		AggregatedEndPoint byte

		Res1 bool
	}
}

func (cmd NmMultiChannelAggregatedMembersGet) CommandClassID() cc.CommandClassID {
	return 0x52
}

func (cmd NmMultiChannelAggregatedMembersGet) CommandID() cc.CommandID {
	return CommandNmMultiChannelAggregatedMembersGet
}

func (cmd NmMultiChannelAggregatedMembersGet) CommandIDString() string {
	return "NM_MULTI_CHANNEL_AGGREGATED_MEMBERS_GET"
}

func (cmd *NmMultiChannelAggregatedMembersGet) UnmarshalBinary(data []byte) error {
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

	cmd.Nodeid = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.AggregatedEndPoint = (payload[i] & 0x7F)

	cmd.Properties1.Res1 = payload[i]&0x80 == 0x80

	i += 1

	return nil
}

func (cmd *NmMultiChannelAggregatedMembersGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SeqNo)

	payload = append(payload, cmd.Nodeid)

	{
		var val byte

		val |= (cmd.Properties1.AggregatedEndPoint) & byte(0x7F)

		if cmd.Properties1.Res1 {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	return
}
