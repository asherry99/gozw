// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package networkmanagementinstallationmaintenancev2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandStatisticsGet cc.CommandID = 0x04

func init() {
	gob.Register(StatisticsGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x67),
		Command:      cc.CommandID(0x04),
		Version:      2,
	}, NewStatisticsGet)
}

func NewStatisticsGet() cc.Command {
	return &StatisticsGet{}
}

// <no value>
type StatisticsGet struct {
	Nodeid byte
}

func (cmd StatisticsGet) CommandClassID() cc.CommandClassID {
	return 0x67
}

func (cmd StatisticsGet) CommandID() cc.CommandID {
	return CommandStatisticsGet
}

func (cmd StatisticsGet) CommandIDString() string {
	return "STATISTICS_GET"
}

func (cmd *StatisticsGet) UnmarshalBinary(data []byte) error {
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

	cmd.Nodeid = payload[i]
	i++

	return nil
}

func (cmd *StatisticsGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.Nodeid)

	return
}
