// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package barrieroperator

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandSignalSupportedReport cc.CommandID = 0x05

func init() {
	gob.Register(SignalSupportedReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x66),
		Command:      cc.CommandID(0x05),
		Version:      1,
	}, NewSignalSupportedReport)
}

func NewSignalSupportedReport() cc.Command {
	return &SignalSupportedReport{}
}

// <no value>
type SignalSupportedReport struct {
	BitMask []byte
}

func (cmd SignalSupportedReport) CommandClassID() cc.CommandClassID {
	return 0x66
}

func (cmd SignalSupportedReport) CommandID() cc.CommandID {
	return CommandSignalSupportedReport
}

func (cmd SignalSupportedReport) CommandIDString() string {
	return "BARRIER_OPERATOR_SIGNAL_SUPPORTED_REPORT"
}

func (cmd *SignalSupportedReport) UnmarshalBinary(data []byte) error {
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

	cmd.BitMask = payload[i:]

	return nil
}

func (cmd *SignalSupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.BitMask...)

	return
}
