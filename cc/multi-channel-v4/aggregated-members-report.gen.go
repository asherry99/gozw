// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package multichannelv4

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandAggregatedMembersReport cc.CommandID = 0x0F

func init() {
	gob.Register(AggregatedMembersReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x60),
		Command:      cc.CommandID(0x0F),
		Version:      4,
	}, NewAggregatedMembersReport)
}

func NewAggregatedMembersReport() cc.Command {
	return &AggregatedMembersReport{}
}

// <no value>
type AggregatedMembersReport struct {
	Properties1 struct {
		AggregatedEndPoint byte
	}

	NumberOfBitMasks byte

	AggregatedMembersBitMask []byte
}

func (cmd AggregatedMembersReport) CommandClassID() cc.CommandClassID {
	return 0x60
}

func (cmd AggregatedMembersReport) CommandID() cc.CommandID {
	return CommandAggregatedMembersReport
}

func (cmd AggregatedMembersReport) CommandIDString() string {
	return "MULTI_CHANNEL_AGGREGATED_MEMBERS_REPORT"
}

func (cmd *AggregatedMembersReport) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.AggregatedEndPoint = (payload[i] & 0x7F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NumberOfBitMasks = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.AggregatedMembersBitMask = payload[i:]

	return nil
}

func (cmd *AggregatedMembersReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.AggregatedEndPoint) & byte(0x7F)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.NumberOfBitMasks)

	payload = append(payload, cmd.AggregatedMembersBitMask...)

	return
}
