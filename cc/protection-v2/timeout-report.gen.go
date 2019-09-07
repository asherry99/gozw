// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package protectionv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandTimeoutReport cc.CommandID = 0x0B

func init() {
	gob.Register(TimeoutReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x75),
		Command:      cc.CommandID(0x0B),
		Version:      2,
	}, NewTimeoutReport)
}

func NewTimeoutReport() cc.Command {
	return &TimeoutReport{}
}

// <no value>
type TimeoutReport struct {
	Timeout byte
}

func (cmd TimeoutReport) CommandClassID() cc.CommandClassID {
	return 0x75
}

func (cmd TimeoutReport) CommandID() cc.CommandID {
	return CommandTimeoutReport
}

func (cmd TimeoutReport) CommandIDString() string {
	return "PROTECTION_TIMEOUT_REPORT"
}

func (cmd *TimeoutReport) UnmarshalBinary(data []byte) error {
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

	cmd.Timeout = payload[i]
	i++

	return nil
}

func (cmd *TimeoutReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.Timeout)

	return
}
