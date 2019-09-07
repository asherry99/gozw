// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package antitheftv2

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandReport cc.CommandID = 0x03

func init() {
	gob.Register(Report{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x5D),
		Command:      cc.CommandID(0x03),
		Version:      2,
	}, NewReport)
}

func NewReport() cc.Command {
	return &Report{}
}

// <no value>
type Report struct {
	AntiTheftProtectionStatus byte

	ManufacturerId uint16

	AntiTheftHintNumberBytes byte

	AntiTheftHintByte []byte
}

func (cmd Report) CommandClassID() cc.CommandClassID {
	return 0x5D
}

func (cmd Report) CommandID() cc.CommandID {
	return CommandReport
}

func (cmd Report) CommandIDString() string {
	return "ANTITHEFT_REPORT"
}

func (cmd *Report) UnmarshalBinary(data []byte) error {
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

	cmd.AntiTheftProtectionStatus = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ManufacturerId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.AntiTheftHintNumberBytes = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	{
		length := (payload[2+2]) & 0xFF
		cmd.AntiTheftHintByte = payload[i : i+int(length)]
		i += int(length)
	}

	return nil
}

func (cmd *Report) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.AntiTheftProtectionStatus)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ManufacturerId)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.AntiTheftHintNumberBytes)

	if cmd.AntiTheftHintByte != nil && len(cmd.AntiTheftHintByte) > 0 {
		payload = append(payload, cmd.AntiTheftHintByte...)
	}

	return
}
