// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package mailbox

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandConfigurationReport cc.CommandID = 0x03

func init() {
	gob.Register(ConfigurationReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x69),
		Command:      cc.CommandID(0x03),
		Version:      1,
	}, NewConfigurationReport)
}

func NewConfigurationReport() cc.Command {
	return &ConfigurationReport{}
}

// <no value>
type ConfigurationReport struct {
	Properties1 struct {
		Mode byte

		SupportedModes byte
	}

	MailboxCapacity uint16

	ForwardingDestinationIpv6Address []byte

	UdpPortNumber uint16
}

func (cmd ConfigurationReport) CommandClassID() cc.CommandClassID {
	return 0x69
}

func (cmd ConfigurationReport) CommandID() cc.CommandID {
	return CommandConfigurationReport
}

func (cmd ConfigurationReport) CommandIDString() string {
	return "MAILBOX_CONFIGURATION_REPORT"
}

func (cmd *ConfigurationReport) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.Mode = (payload[i] & 0x07)

	cmd.Properties1.SupportedModes = (payload[i] & 0x18) >> 3

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MailboxCapacity = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ForwardingDestinationIpv6Address = payload[i : i+16]

	i += 16

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.UdpPortNumber = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}

func (cmd *ConfigurationReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.Mode) & byte(0x07)

		val |= (cmd.Properties1.SupportedModes << byte(3)) & byte(0x18)

		payload = append(payload, val)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.MailboxCapacity)
		payload = append(payload, buf...)
	}

	if paramLen := len(cmd.ForwardingDestinationIpv6Address); paramLen > 16 {
		return nil, errors.New("Length overflow in array parameter ForwardingDestinationIpv6Address")
	}

	payload = append(payload, cmd.ForwardingDestinationIpv6Address...)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.UdpPortNumber)
		payload = append(payload, buf...)
	}

	return
}
