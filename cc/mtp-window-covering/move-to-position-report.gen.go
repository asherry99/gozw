// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package mtpwindowcovering

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandMoveToPositionReport cc.CommandID = 0x03

func init() {
	gob.Register(MoveToPositionReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x51),
		Command:      cc.CommandID(0x03),
		Version:      1,
	}, NewMoveToPositionReport)
}

func NewMoveToPositionReport() cc.Command {
	return &MoveToPositionReport{}
}

// <no value>
type MoveToPositionReport struct {
	Value byte
}

func (cmd MoveToPositionReport) CommandClassID() cc.CommandClassID {
	return 0x51
}

func (cmd MoveToPositionReport) CommandID() cc.CommandID {
	return CommandMoveToPositionReport
}

func (cmd MoveToPositionReport) CommandIDString() string {
	return "MOVE_TO_POSITION_REPORT"
}

func (cmd *MoveToPositionReport) UnmarshalBinary(data []byte) error {
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

	cmd.Value = payload[i]
	i++

	return nil
}

func (cmd *MoveToPositionReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.Value)

	return
}
