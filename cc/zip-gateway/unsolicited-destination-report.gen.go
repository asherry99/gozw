// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package zipgateway

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandUnsolicitedDestinationReport cc.CommandID = 0x0A

func init() {
	gob.Register(UnsolicitedDestinationReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x5F),
		Command:      cc.CommandID(0x0A),
		Version:      1,
	}, NewUnsolicitedDestinationReport)
}

func NewUnsolicitedDestinationReport() cc.Command {
	return &UnsolicitedDestinationReport{}
}

// <no value>
type UnsolicitedDestinationReport struct {
	UnsolicitedIpv6Destination []byte

	UnsolicitedDestinationPort uint16
}

func (cmd UnsolicitedDestinationReport) CommandClassID() cc.CommandClassID {
	return 0x5F
}

func (cmd UnsolicitedDestinationReport) CommandID() cc.CommandID {
	return CommandUnsolicitedDestinationReport
}

func (cmd UnsolicitedDestinationReport) CommandIDString() string {
	return "UNSOLICITED_DESTINATION_REPORT"
}

func (cmd *UnsolicitedDestinationReport) UnmarshalBinary(data []byte) error {
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

	cmd.UnsolicitedIpv6Destination = payload[i : i+16]

	i += 16

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.UnsolicitedDestinationPort = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}

func (cmd *UnsolicitedDestinationReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	if paramLen := len(cmd.UnsolicitedIpv6Destination); paramLen > 16 {
		return nil, errors.New("Length overflow in array parameter UnsolicitedIpv6Destination")
	}

	payload = append(payload, cmd.UnsolicitedIpv6Destination...)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.UnsolicitedDestinationPort)
		payload = append(payload, buf...)
	}

	return
}
