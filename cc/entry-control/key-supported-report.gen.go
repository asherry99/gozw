// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package entrycontrol

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandKeySupportedReport cc.CommandID = 0x03

func init() {
	gob.Register(KeySupportedReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x6F),
		Command:      cc.CommandID(0x03),
		Version:      1,
	}, NewKeySupportedReport)
}

func NewKeySupportedReport() cc.Command {
	return &KeySupportedReport{}
}

// <no value>
type KeySupportedReport struct {
	KeySupportedBitMaskLength byte

	KeySupportedBitMask []byte
}

func (cmd KeySupportedReport) CommandClassID() cc.CommandClassID {
	return 0x6F
}

func (cmd KeySupportedReport) CommandID() cc.CommandID {
	return CommandKeySupportedReport
}

func (cmd KeySupportedReport) CommandIDString() string {
	return "ENTRY_CONTROL_KEY_SUPPORTED_REPORT"
}

func (cmd *KeySupportedReport) UnmarshalBinary(data []byte) error {
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

	cmd.KeySupportedBitMaskLength = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.KeySupportedBitMask = payload[i:]

	return nil
}

func (cmd *KeySupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.KeySupportedBitMaskLength)

	payload = append(payload, cmd.KeySupportedBitMask...)

	return
}
