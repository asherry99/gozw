// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package tarifftblmonitor

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandTariffTblCostGet cc.CommandID = 0x05

func init() {
	gob.Register(TariffTblCostGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x4B),
		Command:      cc.CommandID(0x05),
		Version:      1,
	}, NewTariffTblCostGet)
}

func NewTariffTblCostGet() cc.Command {
	return &TariffTblCostGet{}
}

// <no value>
type TariffTblCostGet struct {
	RateParameterSetId byte

	StartYear uint16

	StartMonth byte

	StartDay byte

	StartHourLocalTime byte

	StartMinuteLocalTime byte

	StopYear uint16

	StopMonth byte

	StopDay byte

	StopHourLocalTime byte

	StopMinuteLocalTime byte
}

func (cmd TariffTblCostGet) CommandClassID() cc.CommandClassID {
	return 0x4B
}

func (cmd TariffTblCostGet) CommandID() cc.CommandID {
	return CommandTariffTblCostGet
}

func (cmd TariffTblCostGet) CommandIDString() string {
	return "TARIFF_TBL_COST_GET"
}

func (cmd *TariffTblCostGet) UnmarshalBinary(data []byte) error {
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

	cmd.RateParameterSetId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartYear = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartMonth = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartDay = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartHourLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartMinuteLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopYear = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopMonth = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopDay = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopHourLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopMinuteLocalTime = payload[i]
	i++

	return nil
}

func (cmd *TariffTblCostGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.RateParameterSetId)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.StartYear)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.StartMonth)

	payload = append(payload, cmd.StartDay)

	payload = append(payload, cmd.StartHourLocalTime)

	payload = append(payload, cmd.StartMinuteLocalTime)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.StopYear)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.StopMonth)

	payload = append(payload, cmd.StopDay)

	payload = append(payload, cmd.StopHourLocalTime)

	payload = append(payload, cmd.StopMinuteLocalTime)

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}
