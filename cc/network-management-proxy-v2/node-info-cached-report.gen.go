// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package networkmanagementproxyv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandNodeInfoCachedReport cc.CommandID = 0x04

func init() {
	gob.Register(NodeInfoCachedReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x52),
		Command:      cc.CommandID(0x04),
		Version:      2,
	}, NewNodeInfoCachedReport)
}

func NewNodeInfoCachedReport() cc.Command {
	return &NodeInfoCachedReport{}
}

// <no value>
type NodeInfoCachedReport struct {
	SeqNo byte

	Properties1 struct {
		Age byte

		Status byte
	}

	Properties2 struct {
		ZWaveProtocolSpecificPart1 byte

		Listening bool
	}

	Properties3 struct {
		ZWaveProtocolSpecificPart2 byte

		Opt bool
	}

	GrantedKeys byte

	BasicDeviceClass byte

	GenericDeviceClass byte

	SpecificDeviceClass byte

	CommandClass []byte
}

func (cmd NodeInfoCachedReport) CommandClassID() cc.CommandClassID {
	return 0x52
}

func (cmd NodeInfoCachedReport) CommandID() cc.CommandID {
	return CommandNodeInfoCachedReport
}

func (cmd NodeInfoCachedReport) CommandIDString() string {
	return "NODE_INFO_CACHED_REPORT"
}

func (cmd *NodeInfoCachedReport) UnmarshalBinary(data []byte) error {
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

	cmd.SeqNo = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.Age = (payload[i] & 0x0F)

	cmd.Properties1.Status = (payload[i] & 0xF0) >> 4

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties2.ZWaveProtocolSpecificPart1 = (payload[i] & 0x7F)

	cmd.Properties2.Listening = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties3.ZWaveProtocolSpecificPart2 = (payload[i] & 0x7F)

	cmd.Properties3.Opt = payload[i]&0x80 == 0x80

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GrantedKeys = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.BasicDeviceClass = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GenericDeviceClass = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SpecificDeviceClass = payload[i]
	i++

	if len(payload) <= i {
		return nil
	}

	cmd.CommandClass = payload[i:]

	return nil
}

func (cmd *NodeInfoCachedReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SeqNo)

	{
		var val byte

		val |= (cmd.Properties1.Age) & byte(0x0F)

		val |= (cmd.Properties1.Status << byte(4)) & byte(0xF0)

		payload = append(payload, val)
	}

	{
		var val byte

		val |= (cmd.Properties2.ZWaveProtocolSpecificPart1) & byte(0x7F)

		if cmd.Properties2.Listening {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	{
		var val byte

		val |= (cmd.Properties3.ZWaveProtocolSpecificPart2) & byte(0x7F)

		if cmd.Properties3.Opt {
			val |= byte(0x80) // flip bits on
		} else {
			val &= ^byte(0x80) // flip bits off
		}

		payload = append(payload, val)
	}

	payload = append(payload, cmd.GrantedKeys)

	payload = append(payload, cmd.BasicDeviceClass)

	payload = append(payload, cmd.GenericDeviceClass)

	payload = append(payload, cmd.SpecificDeviceClass)

	payload = append(payload, cmd.CommandClass...)

	return
}
