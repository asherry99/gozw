// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package networkmanagementinstallationmaintenancev2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandRssiReport cc.CommandID = 0x08

func init() {
	gob.Register(RssiReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x67),
		Command:      cc.CommandID(0x08),
		Version:      2,
	}, NewRssiReport)
}

func NewRssiReport() cc.Command {
	return &RssiReport{}
}

// <no value>
type RssiReport struct {
	Channel1Rssi byte

	Channel2Rssi byte

	Channel3Rssi byte
}

func (cmd RssiReport) CommandClassID() cc.CommandClassID {
	return 0x67
}

func (cmd RssiReport) CommandID() cc.CommandID {
	return CommandRssiReport
}

func (cmd RssiReport) CommandIDString() string {
	return "RSSI_REPORT"
}

func (cmd *RssiReport) UnmarshalBinary(data []byte) error {
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

	cmd.Channel1Rssi = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Channel2Rssi = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Channel3Rssi = payload[i]
	i++

	return nil
}

func (cmd *RssiReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.Channel1Rssi)

	payload = append(payload, cmd.Channel2Rssi)

	payload = append(payload, cmd.Channel3Rssi)

	return
}
