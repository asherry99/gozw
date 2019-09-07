// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package dcpconfig

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandDcpListSet cc.CommandID = 0x03

func init() {
	gob.Register(DcpListSet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x3A),
		Command:      cc.CommandID(0x03),
		Version:      1,
	}, NewDcpListSet)
}

func NewDcpListSet() cc.Command {
	return &DcpListSet{}
}

// <no value>
type DcpListSet struct {
	Year uint16

	Month byte

	Day byte

	HourLocalTime byte

	MinuteLocalTime byte

	SecondLocalTime byte

	DcpRateId byte

	Properties1 struct {
		NumberOfDc byte
	}

	StartYear uint16

	StartMonth byte

	StartDay byte

	StartHourLocalTime byte

	StartMinuteLocalTime byte

	StartSecondLocalTime byte

	DurationHourTime byte

	DurationMinuteTime byte

	DurationSecondTime byte

	EventPriority byte

	LoadShedding byte

	StartAssociationGroup byte

	StopAssociationGroup byte

	RandomizationInterval byte

	Vg1 []DcpListSetVg1
}

type DcpListSetVg1 struct {
	GenericDeviceClass byte

	SpecificDeviceClass byte
}

func (cmd DcpListSet) CommandClassID() cc.CommandClassID {
	return 0x3A
}

func (cmd DcpListSet) CommandID() cc.CommandID {
	return CommandDcpListSet
}

func (cmd DcpListSet) CommandIDString() string {
	return "DCP_LIST_SET"
}

func (cmd *DcpListSet) UnmarshalBinary(data []byte) error {
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

	cmd.Year = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Month = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Day = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.HourLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MinuteLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SecondLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DcpRateId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.NumberOfDc = (payload[i] & 0x03)

	i += 1

	for i < int(cmd.Properties1.NumberOfDc) {

		vg1 := DcpListSetVg1{}

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		genericDeviceClass := payload[i]
		i++

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		specificDeviceClass := payload[i]
		i++

		vg1.GenericDeviceClass = genericDeviceClass

		vg1.SpecificDeviceClass = specificDeviceClass

		cmd.Vg1 = append(cmd.Vg1, vg1)
	}

	cmd.StartYear = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.StartMonth = payload[i]
	i++

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.StartDay = payload[i]
	i++

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.StartHourLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.StartMinuteLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.StartSecondLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.DurationHourTime = payload[i]
	i++

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.DurationMinuteTime = payload[i]
	i++

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.DurationSecondTime = payload[i]
	i++

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.EventPriority = payload[i]
	i++

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.LoadShedding = payload[i]
	i++

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.StartAssociationGroup = payload[i]
	i++

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.StopAssociationGroup = payload[i]
	i++

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.RandomizationInterval = payload[i]
	i++

	return nil
}

func (cmd *DcpListSet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.Year)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.Month)

	payload = append(payload, cmd.Day)

	payload = append(payload, cmd.HourLocalTime)

	payload = append(payload, cmd.MinuteLocalTime)

	payload = append(payload, cmd.SecondLocalTime)

	payload = append(payload, cmd.DcpRateId)

	{
		var val byte

		val |= (cmd.Properties1.NumberOfDc) & byte(0x03)

		payload = append(payload, val)
	}

	for _, vg := range cmd.Vg1 {

		payload = append(payload, vg.GenericDeviceClass)

		payload = append(payload, vg.SpecificDeviceClass)

	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.StartYear)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.StartMonth)

	payload = append(payload, cmd.StartDay)

	payload = append(payload, cmd.StartHourLocalTime)

	payload = append(payload, cmd.StartMinuteLocalTime)

	payload = append(payload, cmd.StartSecondLocalTime)

	payload = append(payload, cmd.DurationHourTime)

	payload = append(payload, cmd.DurationMinuteTime)

	payload = append(payload, cmd.DurationSecondTime)

	payload = append(payload, cmd.EventPriority)

	payload = append(payload, cmd.LoadShedding)

	payload = append(payload, cmd.StartAssociationGroup)

	payload = append(payload, cmd.StopAssociationGroup)

	payload = append(payload, cmd.RandomizationInterval)

	return
}
