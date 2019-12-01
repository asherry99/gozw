// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package ratetblmonitor

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandRateTblActiveRateReport cc.CommandID = 0x06

func init() {
	gob.Register(RateTblActiveRateReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x49),
		Command:      cc.CommandID(0x06),
		Version:      1,
	}, NewRateTblActiveRateReport)
}

func NewRateTblActiveRateReport() cc.Command {
	return &RateTblActiveRateReport{}
}

// <no value>
type RateTblActiveRateReport struct {
	RateParameterSetId byte
}

func (cmd RateTblActiveRateReport) CommandClassID() cc.CommandClassID {
	return 0x49
}

func (cmd RateTblActiveRateReport) CommandID() cc.CommandID {
	return CommandRateTblActiveRateReport
}

func (cmd RateTblActiveRateReport) CommandIDString() string {
	return "RATE_TBL_ACTIVE_RATE_REPORT"
}

func (cmd *RateTblActiveRateReport) UnmarshalBinary(data []byte) error {
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

	cmd.RateParameterSetId = payload[i]
	i++

	return nil
}

func (cmd *RateTblActiveRateReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.RateParameterSetId)

	return
}