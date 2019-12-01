// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package chimneyfan

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandModeReport cc.CommandID = 0x18

func init() {
	gob.Register(ModeReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x2A),
		Command:      cc.CommandID(0x18),
		Version:      1,
	}, NewModeReport)
}

func NewModeReport() cc.Command {
	return &ModeReport{}
}

// <no value>
type ModeReport struct {
	Mode byte
}

func (cmd ModeReport) CommandClassID() cc.CommandClassID {
	return 0x2A
}

func (cmd ModeReport) CommandID() cc.CommandID {
	return CommandModeReport
}

func (cmd ModeReport) CommandIDString() string {
	return "CHIMNEY_FAN_MODE_REPORT"
}

func (cmd *ModeReport) UnmarshalBinary(data []byte) error {
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

	cmd.Mode = payload[i]
	i++

	return nil
}

func (cmd *ModeReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.Mode)

	return
}