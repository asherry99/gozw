// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package multichannelv4

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandMultiInstanceReport cc.CommandID = 0x05

func init() {
	gob.Register(MultiInstanceReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x60),
		Command:      cc.CommandID(0x05),
		Version:      4,
	}, NewMultiInstanceReport)
}

func NewMultiInstanceReport() cc.Command {
	return &MultiInstanceReport{}
}

// <no value>
type MultiInstanceReport struct {
	CommandClass byte

	Properties1 struct {
		Instances byte
	}
}

func (cmd MultiInstanceReport) CommandClassID() cc.CommandClassID {
	return 0x60
}

func (cmd MultiInstanceReport) CommandID() cc.CommandID {
	return CommandMultiInstanceReport
}

func (cmd MultiInstanceReport) CommandIDString() string {
	return "MULTI_INSTANCE_REPORT"
}

func (cmd *MultiInstanceReport) UnmarshalBinary(data []byte) error {
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

	cmd.CommandClass = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.Instances = (payload[i] & 0x7F)

	i += 1

	return nil
}

func (cmd *MultiInstanceReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.CommandClass)

	{
		var val byte

		val |= (cmd.Properties1.Instances) & byte(0x7F)

		payload = append(payload, val)
	}

	return
}
