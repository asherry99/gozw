// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package prepayment

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandBalanceReport cc.CommandID = 0x02

func init() {
	gob.Register(BalanceReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x3F),
		Command:      cc.CommandID(0x02),
		Version:      1,
	}, NewBalanceReport)
}

func NewBalanceReport() cc.Command {
	return &BalanceReport{}
}

// <no value>
type BalanceReport struct {
	Properties1 struct {
		MeterType byte

		BalanceType byte
	}

	Properties2 struct {
		Scale byte

		BalancePrecision byte
	}

	BalanceValue uint32

	Properties3 struct {
		DebtPrecision byte
	}

	Debt uint32

	Properties4 struct {
		EmerCreditPrecision byte
	}

	EmerCredit uint32

	Currency uint32

	DebtRecoveryPercentage byte
}

func (cmd BalanceReport) CommandClassID() cc.CommandClassID {
	return 0x3F
}

func (cmd BalanceReport) CommandID() cc.CommandID {
	return CommandBalanceReport
}

func (cmd BalanceReport) CommandIDString() string {
	return "PREPAYMENT_BALANCE_REPORT"
}

func (cmd *BalanceReport) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.MeterType = (payload[i] & 0x3F)

	cmd.Properties1.BalanceType = (payload[i] & 0xC0) >> 6

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties2.Scale = (payload[i] & 0x1F)

	cmd.Properties2.BalancePrecision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.BalanceValue = binary.BigEndian.Uint32(payload[i : i+4])
	i += 4

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties3.DebtPrecision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Debt = binary.BigEndian.Uint32(payload[i : i+4])
	i += 4

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties4.EmerCreditPrecision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.EmerCredit = binary.BigEndian.Uint32(payload[i : i+4])
	i += 4

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Currency = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DebtRecoveryPercentage = payload[i]
	i++

	return nil
}

func (cmd *BalanceReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.MeterType) & byte(0x3F)

		val |= (cmd.Properties1.BalanceType << byte(6)) & byte(0xC0)

		payload = append(payload, val)
	}

	{
		var val byte

		val |= (cmd.Properties2.Scale) & byte(0x1F)

		val |= (cmd.Properties2.BalancePrecision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.BalanceValue)
		payload = append(payload, buf...)
	}

	{
		var val byte

		val |= (cmd.Properties3.DebtPrecision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.Debt)
		payload = append(payload, buf...)
	}

	{
		var val byte

		val |= (cmd.Properties4.EmerCreditPrecision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.EmerCredit)
		payload = append(payload, buf...)
	}

	{
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, cmd.Currency)
		if buf[0] != 0 {
			return nil, errors.New("BIT_24 value overflow")
		}
		payload = append(payload, buf[1:4]...)
	}

	payload = append(payload, cmd.DebtRecoveryPercentage)

	return
}