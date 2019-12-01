// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package dcpmonitor

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandDcpEventStatusGet cc.CommandID = 0x03

func init() {
	gob.Register(DcpEventStatusGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x3B),
		Command:      cc.CommandID(0x03),
		Version:      1,
	}, NewDcpEventStatusGet)
}

func NewDcpEventStatusGet() cc.Command {
	return &DcpEventStatusGet{}
}

// <no value>
type DcpEventStatusGet struct {
	Year uint16

	Month byte

	Day byte

	HourLocalTime byte

	MinuteLocalTime byte

	SecondLocalTime byte
}

func (cmd DcpEventStatusGet) CommandClassID() cc.CommandClassID {
	return 0x3B
}

func (cmd DcpEventStatusGet) CommandID() cc.CommandID {
	return CommandDcpEventStatusGet
}

func (cmd DcpEventStatusGet) CommandIDString() string {
	return "DCP_EVENT_STATUS_GET"
}

func (cmd *DcpEventStatusGet) UnmarshalBinary(data []byte) error {
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

	return nil
}

func (cmd *DcpEventStatusGet) MarshalBinary() (payload []byte, err error) {
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

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}