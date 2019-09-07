// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package configurationv2

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandBulkGet cc.CommandID = 0x08

func init() {
	gob.Register(BulkGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x70),
		Command:      cc.CommandID(0x08),
		Version:      2,
	}, NewBulkGet)
}

func NewBulkGet() cc.Command {
	return &BulkGet{}
}

// <no value>
type BulkGet struct {
	ParameterOffset uint16

	NumberOfParameters byte
}

func (cmd BulkGet) CommandClassID() cc.CommandClassID {
	return 0x70
}

func (cmd BulkGet) CommandID() cc.CommandID {
	return CommandBulkGet
}

func (cmd BulkGet) CommandIDString() string {
	return "CONFIGURATION_BULK_GET"
}

func (cmd *BulkGet) UnmarshalBinary(data []byte) error {
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

	cmd.ParameterOffset = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NumberOfParameters = payload[i]
	i++

	return nil
}

func (cmd *BulkGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ParameterOffset)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.NumberOfParameters)

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}
