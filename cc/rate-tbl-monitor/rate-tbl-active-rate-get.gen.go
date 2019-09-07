// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package ratetblmonitor

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandRateTblActiveRateGet cc.CommandID = 0x05

func init() {
	gob.Register(RateTblActiveRateGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x49),
		Command:      cc.CommandID(0x05),
		Version:      1,
	}, NewRateTblActiveRateGet)
}

func NewRateTblActiveRateGet() cc.Command {
	return &RateTblActiveRateGet{}
}

// <no value>
type RateTblActiveRateGet struct {
}

func (cmd RateTblActiveRateGet) CommandClassID() cc.CommandClassID {
	return 0x49
}

func (cmd RateTblActiveRateGet) CommandID() cc.CommandID {
	return CommandRateTblActiveRateGet
}

func (cmd RateTblActiveRateGet) CommandIDString() string {
	return "RATE_TBL_ACTIVE_RATE_GET"
}

func (cmd *RateTblActiveRateGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *RateTblActiveRateGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}
