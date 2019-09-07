// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package irrigation

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandValveConfigReport cc.CommandID = 0x0C

func init() {
	gob.Register(ValveConfigReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x6B),
		Command:      cc.CommandID(0x0C),
		Version:      1,
	}, NewValveConfigReport)
}

func NewValveConfigReport() cc.Command {
	return &ValveConfigReport{}
}

// <no value>
type ValveConfigReport struct {
	Properties1 struct {
		MasterValve bool
	}

	ValveId byte

	NominalCurrentHighThreshold byte

	NominalCurrentLowThreshold byte

	Properties2 struct {
		MaximumFlowSize byte

		MaximumFlowScale byte

		MaximumFlowPrecision byte
	}

	MaximumFlowValue []byte

	Properties3 struct {
		FlowHighThresholdSize byte

		FlowHighThresholdScale byte

		FlowHighThresholdPrecision byte
	}

	FlowHighThresholdValue []byte

	Properties4 struct {
		FlowLowThresholdSize byte

		FlowLowThresholdScale byte

		FlowLowThresholdPrecision byte
	}

	FlowLowThresholdValue []byte

	SensorUsage []byte
}

func (cmd ValveConfigReport) CommandClassID() cc.CommandClassID {
	return 0x6B
}

func (cmd ValveConfigReport) CommandID() cc.CommandID {
	return CommandValveConfigReport
}

func (cmd ValveConfigReport) CommandIDString() string {
	return "IRRIGATION_VALVE_CONFIG_REPORT"
}

func (cmd *ValveConfigReport) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.MasterValve = payload[i]&0x01 == 0x01

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ValveId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NominalCurrentHighThreshold = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NominalCurrentLowThreshold = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties2.MaximumFlowSize = (payload[i] & 0x07)

	cmd.Properties2.MaximumFlowScale = (payload[i] & 0x18) >> 3

	cmd.Properties2.MaximumFlowPrecision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return nil // field is optional
	}

	{
		length := (payload[4+2]) & 0x07
		cmd.MaximumFlowValue = payload[i : i+int(length)]
		i += int(length)
	}

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.Properties3.FlowHighThresholdSize = (payload[i] & 0x07)

	cmd.Properties3.FlowHighThresholdScale = (payload[i] & 0x18) >> 3

	cmd.Properties3.FlowHighThresholdPrecision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return nil // field is optional
	}

	{
		length := (payload[6+2]) & 0x07
		cmd.FlowHighThresholdValue = payload[i : i+int(length)]
		i += int(length)
	}

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.Properties4.FlowLowThresholdSize = (payload[i] & 0x07)

	cmd.Properties4.FlowLowThresholdScale = (payload[i] & 0x18) >> 3

	cmd.Properties4.FlowLowThresholdPrecision = (payload[i] & 0xE0) >> 5

	i += 1

	if len(payload) <= i {
		return nil // field is optional
	}

	{
		length := (payload[8+2]) & 0x07
		cmd.FlowLowThresholdValue = payload[i : i+int(length)]
		i += int(length)
	}

	if len(payload) <= i {
		return nil // field is optional
	}

	cmd.SensorUsage = payload[i:]

	return nil
}

func (cmd *ValveConfigReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		if cmd.Properties1.MasterValve {
			val |= byte(0x01) // flip bits on
		} else {
			val &= ^byte(0x01) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.ValveId)

	payload = append(payload, cmd.NominalCurrentHighThreshold)

	payload = append(payload, cmd.NominalCurrentLowThreshold)

	{
		var val byte

		val |= (cmd.Properties2.MaximumFlowSize) & byte(0x07)

		val |= (cmd.Properties2.MaximumFlowScale << byte(3)) & byte(0x18)

		val |= (cmd.Properties2.MaximumFlowPrecision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	if cmd.MaximumFlowValue != nil && len(cmd.MaximumFlowValue) > 0 {
		payload = append(payload, cmd.MaximumFlowValue...)
	}

	{
		var val byte

		val |= (cmd.Properties3.FlowHighThresholdSize) & byte(0x07)

		val |= (cmd.Properties3.FlowHighThresholdScale << byte(3)) & byte(0x18)

		val |= (cmd.Properties3.FlowHighThresholdPrecision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	if cmd.FlowHighThresholdValue != nil && len(cmd.FlowHighThresholdValue) > 0 {
		payload = append(payload, cmd.FlowHighThresholdValue...)
	}

	{
		var val byte

		val |= (cmd.Properties4.FlowLowThresholdSize) & byte(0x07)

		val |= (cmd.Properties4.FlowLowThresholdScale << byte(3)) & byte(0x18)

		val |= (cmd.Properties4.FlowLowThresholdPrecision << byte(5)) & byte(0xE0)

		payload = append(payload, val)
	}

	if cmd.FlowLowThresholdValue != nil && len(cmd.FlowLowThresholdValue) > 0 {
		payload = append(payload, cmd.FlowLowThresholdValue...)
	}

	payload = append(payload, cmd.SensorUsage...)

	return
}
