// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package zwaveplusinfov2

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandReport cc.CommandID = 0x02

func init() {
	gob.Register(Report{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x5E),
		Command:      cc.CommandID(0x02),
		Version:      2,
	}, NewReport)
}

func NewReport() cc.Command {
	return &Report{}
}

// <no value>
type Report struct {
	ZWaveVersion byte

	RoleType byte

	NodeType byte

	InstallerIconType uint16

	UserIconType uint16
}

func (cmd Report) CommandClassID() cc.CommandClassID {
	return 0x5E
}

func (cmd Report) CommandID() cc.CommandID {
	return CommandReport
}

func (cmd Report) CommandIDString() string {
	return "ZWAVEPLUS_INFO_REPORT"
}

func (cmd *Report) UnmarshalBinary(data []byte) error {
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

	cmd.ZWaveVersion = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.RoleType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NodeType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.InstallerIconType = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.UserIconType = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}

func (cmd *Report) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.ZWaveVersion)

	payload = append(payload, cmd.RoleType)

	payload = append(payload, cmd.NodeType)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.InstallerIconType)
		payload = append(payload, buf...)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.UserIconType)
		payload = append(payload, buf...)
	}

	return
}
