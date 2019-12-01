// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package schedule

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandStateReport cc.CommandID = 0x09

func init() {
	gob.Register(StateReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x53),
		Command:      cc.CommandID(0x09),
		Version:      1,
	}, NewStateReport)
}

func NewStateReport() cc.Command {
	return &StateReport{}
}

// <no value>
type StateReport struct {
	NumberOfSupportedScheduleId byte

	Properties1 struct {
		ReportsToFollow byte

		Override bool
	}

	Vg1 []StateReportVg1
}

type StateReportVg1 struct {
	Properties2 struct {
		ActiveId1 byte

		ActiveId2 byte
	}
}

func (cmd StateReport) CommandClassID() cc.CommandClassID {
	return 0x53
}

func (cmd StateReport) CommandID() cc.CommandID {
	return CommandStateReport
}

func (cmd StateReport) CommandIDString() string {
	return "SCHEDULE_STATE_REPORT"
}

func (cmd *StateReport) UnmarshalBinary(data []byte) error {
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

	cmd.NumberOfSupportedScheduleId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.ReportsToFollow = (payload[i] & 0xFE) >> 1

	cmd.Properties1.Override = payload[i]&0x01 == 0x01

	i += 1

	for i < len(payload) {

		vg1 := StateReportVg1{}

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		vg1.Properties2.ActiveId1 = (payload[i] & 0x0F)

		vg1.Properties2.ActiveId2 = (payload[i] & 0xF0) >> 4

		i += 1

		// struct byte fields are assigned to the variant group when computed

		cmd.Vg1 = append(cmd.Vg1, vg1)
	}

	return nil
}

func (cmd *StateReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.NumberOfSupportedScheduleId)

	{
		var val byte

		val |= (cmd.Properties1.ReportsToFollow << byte(1)) & byte(0xFE)

		if cmd.Properties1.Override {
			val |= byte(0x01) // flip bits on
		} else {
			val &= ^byte(0x01) // flip bits off
		}

		payload = append(payload, val)
	}

	for _, vg := range cmd.Vg1 {

		{
			var val byte

			val |= (vg.Properties2.ActiveId1) & byte(0x0F)

			val |= (vg.Properties2.ActiveId2 << byte(4)) & byte(0xF0)

			payload = append(payload, val)
		}

	}

	return
}