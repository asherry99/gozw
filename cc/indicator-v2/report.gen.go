// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package indicatorv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandReport cc.CommandID = 0x03

func init() {
	gob.Register(Report{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x87),
		Command:      cc.CommandID(0x03),
		Version:      2,
	}, NewReport)
}

func NewReport() cc.Command {
	return &Report{}
}

// <no value>
type Report struct {
	Indicator0Value byte

	Properties1 struct {
		IndicatorObjectCount byte
	}

	Vg1 []ReportVg1
}

type ReportVg1 struct {
	IndicatorId byte

	PropertyId byte

	Value byte
}

func (cmd Report) CommandClassID() cc.CommandClassID {
	return 0x87
}

func (cmd Report) CommandID() cc.CommandID {
	return CommandReport
}

func (cmd Report) CommandIDString() string {
	return "INDICATOR_REPORT"
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

	cmd.Indicator0Value = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.IndicatorObjectCount = (payload[i] & 0x1F)

	i += 1

	for i < len(payload) {

		vg1 := ReportVg1{}

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		indicatorId := payload[i]
		i++

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		propertyId := payload[i]
		i++

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		value := payload[i]
		i++

		vg1.IndicatorId = indicatorId

		vg1.PropertyId = propertyId

		vg1.Value = value

		cmd.Vg1 = append(cmd.Vg1, vg1)
	}

	return nil
}

func (cmd *Report) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.Indicator0Value)

	{
		var val byte

		val |= (cmd.Properties1.IndicatorObjectCount) & byte(0x1F)

		payload = append(payload, val)
	}

	for _, vg := range cmd.Vg1 {

		payload = append(payload, vg.IndicatorId)

		payload = append(payload, vg.PropertyId)

		payload = append(payload, vg.Value)

	}

	return
}
