// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package ratetblmonitor

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandRateTblSupportedReport cc.CommandID = 0x02

func init() {
	gob.Register(RateTblSupportedReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x49),
		Command:      cc.CommandID(0x02),
		Version:      1,
	}, NewRateTblSupportedReport)
}

func NewRateTblSupportedReport() cc.Command {
	return &RateTblSupportedReport{}
}

// <no value>
type RateTblSupportedReport struct {
	RatesSupported byte

	ParameterSetSupportedBitMask uint16
}

func (cmd RateTblSupportedReport) CommandClassID() cc.CommandClassID {
	return 0x49
}

func (cmd RateTblSupportedReport) CommandID() cc.CommandID {
	return CommandRateTblSupportedReport
}

func (cmd RateTblSupportedReport) CommandIDString() string {
	return "RATE_TBL_SUPPORTED_REPORT"
}

func (cmd *RateTblSupportedReport) UnmarshalBinary(data []byte) error {
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

	cmd.RatesSupported = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ParameterSetSupportedBitMask = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}

func (cmd *RateTblSupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.RatesSupported)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ParameterSetSupportedBitMask)
		payload = append(payload, buf...)
	}

	return
}
