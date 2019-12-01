// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package networkmanagementinclusionv3

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandS2AdvancedJoinModeReport cc.CommandID = 0x18

func init() {
	gob.Register(S2AdvancedJoinModeReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x34),
		Command:      cc.CommandID(0x18),
		Version:      3,
	}, NewS2AdvancedJoinModeReport)
}

func NewS2AdvancedJoinModeReport() cc.Command {
	return &S2AdvancedJoinModeReport{}
}

// <no value>
type S2AdvancedJoinModeReport struct {
	SeqNo byte

	S2AdvancedJoinMode byte
}

func (cmd S2AdvancedJoinModeReport) CommandClassID() cc.CommandClassID {
	return 0x34
}

func (cmd S2AdvancedJoinModeReport) CommandID() cc.CommandID {
	return CommandS2AdvancedJoinModeReport
}

func (cmd S2AdvancedJoinModeReport) CommandIDString() string {
	return "S2_ADVANCED_JOIN_MODE_REPORT"
}

func (cmd *S2AdvancedJoinModeReport) UnmarshalBinary(data []byte) error {
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

	cmd.S2AdvancedJoinMode = payload[i]
	i++

	return nil
}

func (cmd *S2AdvancedJoinModeReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SeqNo)

	payload = append(payload, cmd.S2AdvancedJoinMode)

	return
}