// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package centralscenev2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandSupportedReport cc.CommandID = 0x02

func init() {
	gob.Register(SupportedReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x5B),
		Command:      cc.CommandID(0x02),
		Version:      2,
	}, NewSupportedReport)
}

func NewSupportedReport() cc.Command {
	return &SupportedReport{}
}

// <no value>
type SupportedReport struct {
	SupportedScenes byte

	Properties1 struct {
		NumberOfBitMaskBytes byte

		Identical bool
	}

	Vg1 []SupportedReportVg1
}

type SupportedReportVg1 struct {
	SupportedKeyAttributesForScene []byte
}

func (cmd SupportedReport) CommandClassID() cc.CommandClassID {
	return 0x5B
}

func (cmd SupportedReport) CommandID() cc.CommandID {
	return CommandSupportedReport
}

func (cmd SupportedReport) CommandIDString() string {
	return "CENTRAL_SCENE_SUPPORTED_REPORT"
}

func (cmd *SupportedReport) UnmarshalBinary(data []byte) error {
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

	cmd.SupportedScenes = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.NumberOfBitMaskBytes = (payload[i] & 0x06) >> 1

	cmd.Properties1.Identical = payload[i]&0x01 == 0x01

	i += 1

	for i < len(payload) {

		vg1 := SupportedReportVg1{}

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		supportedKeyAttributesForScene := payload[i:]

		vg1.SupportedKeyAttributesForScene = supportedKeyAttributesForScene

		cmd.Vg1 = append(cmd.Vg1, vg1)
	}

	return nil
}

func (cmd *SupportedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SupportedScenes)

	{
		var val byte

		val |= (cmd.Properties1.NumberOfBitMaskBytes << byte(1)) & byte(0x06)

		if cmd.Properties1.Identical {
			val |= byte(0x01) // flip bits on
		} else {
			val &= ^byte(0x01) // flip bits off
		}

		payload = append(payload, val)
	}

	for _, vg := range cmd.Vg1 {

		payload = append(payload, vg.SupportedKeyAttributesForScene...)

	}

	return
}
