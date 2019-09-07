// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package schedulev3

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandSet cc.CommandID = 0x03

func init() {
	gob.Register(Set{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x53),
		Command:      cc.CommandID(0x03),
		Version:      3,
	}, NewSet)
}

func NewSet() cc.Command {
	return &Set{}
}

// <no value>
type Set struct {
	ScheduleId byte

	ScheduleIdBlock byte

	StartYear byte

	Properties1 struct {
		StartMonth byte

		RecurrenceOffset byte
	}

	Properties2 struct {
		StartDayOfMonth byte

		RecurrenceMode byte
	}

	Properties3 struct {
		StartWeekday byte
	}

	Properties4 struct {
		StartHour byte

		DurationType byte
	}

	Properties5 struct {
		StartMinute byte

		Relative bool
	}

	DurationByte uint16

	ReportsToFollow byte

	NumberOfCmdToFollow byte

	Vg1 []SetVg1
}

type SetVg1 struct {
	CmdLength byte

	CmdByte []byte
}

func (cmd Set) CommandClassID() cc.CommandClassID {
	return 0x53
}

func (cmd Set) CommandID() cc.CommandID {
	return CommandSet
}

func (cmd Set) CommandIDString() string {
	return "COMMAND_SCHEDULE_SET"
}

func (cmd *Set) UnmarshalBinary(data []byte) error {
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

	cmd.ScheduleId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ScheduleIdBlock = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartYear = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.StartMonth = (payload[i] & 0x0F)

	cmd.Properties1.RecurrenceOffset = (payload[i] & 0xF0) >> 4

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties2.StartDayOfMonth = (payload[i] & 0x1F)

	cmd.Properties2.RecurrenceMode = (payload[i] & 0x60) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties3.StartWeekday = (payload[i] & 0x7F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties4.StartHour = (payload[i] & 0x1F)

	cmd.Properties4.DurationType = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties5.StartMinute = (payload[i] & 0x3F)

	cmd.Properties5.Relative = payload[i]&0x40 == 0x40

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DurationByte = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ReportsToFollow = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NumberOfCmdToFollow = payload[i]
	i++

	for i < len(payload) {

		vg1 := SetVg1{}

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		cmdLength := payload[i]
		i++

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		length := (payload[0+2]) & 0xFF
		cmdByte := payload[i : i+int(length)]
		i += int(length)

		vg1.CmdLength = cmdLength

		vg1.CmdByte = cmdByte

		cmd.Vg1 = append(cmd.Vg1, vg1)
	}

	return nil
}

func (cmd *Set) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.ScheduleId)

	payload = append(payload, cmd.ScheduleIdBlock)

	payload = append(payload, cmd.StartYear)

	{
		var val byte

		val |= (cmd.Properties1.StartMonth) & byte(0x0F)

		val |= (cmd.Properties1.RecurrenceOffset << byte(4)) & byte(0xF0)

		payload = append(payload, val)
	}

	{
		var val byte

		val |= (cmd.Properties2.StartDayOfMonth) & byte(0x1F)

		val |= (cmd.Properties2.RecurrenceMode << byte(5)) & byte(0x60)

		payload = append(payload, val)
	}

	{
		var val byte

		val |= (cmd.Properties3.StartWeekday) & byte(0x7F)

		payload = append(payload, val)
	}

	{
		var val byte

		val |= (cmd.Properties4.StartHour) & byte(0x1F)

		val |= (cmd.Properties4.DurationType << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	{
		var val byte

		val |= (cmd.Properties5.StartMinute) & byte(0x3F)

		if cmd.Properties5.Relative {
			val |= byte(0x40) // flip bits on
		} else {
			val &= ^byte(0x40) // flip bits off
		}

		payload = append(payload, val)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.DurationByte)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.ReportsToFollow)

	payload = append(payload, cmd.NumberOfCmdToFollow)

	for _, vg := range cmd.Vg1 {

		payload = append(payload, vg.CmdLength)

		if vg.CmdByte != nil && len(vg.CmdByte) > 0 {
			payload = append(payload, vg.CmdByte...)
		}

	}

	return
}
