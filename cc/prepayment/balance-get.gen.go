// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package prepayment

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandBalanceGet cc.CommandID = 0x01

func init() {
	gob.Register(BalanceGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x3F),
		Command:      cc.CommandID(0x01),
		Version:      1,
	}, NewBalanceGet)
}

func NewBalanceGet() cc.Command {
	return &BalanceGet{}
}

// <no value>
type BalanceGet struct {
	Properties1 struct {
		BalanceType byte
	}
}

func (cmd BalanceGet) CommandClassID() cc.CommandClassID {
	return 0x3F
}

func (cmd BalanceGet) CommandID() cc.CommandID {
	return CommandBalanceGet
}

func (cmd BalanceGet) CommandIDString() string {
	return "PREPAYMENT_BALANCE_GET"
}

func (cmd *BalanceGet) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.BalanceType = (payload[i] & 0xC0) >> 6

	i += 1

	return nil
}

func (cmd *BalanceGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.BalanceType << byte(6)) & byte(0xC0)

		payload = append(payload, val)
	}

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}
