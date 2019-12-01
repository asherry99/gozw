// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package zipgateway

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandApplicationNodeInfoReport cc.CommandID = 0x0D

func init() {
	gob.Register(ApplicationNodeInfoReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x5F),
		Command:      cc.CommandID(0x0D),
		Version:      1,
	}, NewApplicationNodeInfoReport)
}

func NewApplicationNodeInfoReport() cc.Command {
	return &ApplicationNodeInfoReport{}
}

// <no value>
type ApplicationNodeInfoReport struct {
	NonSecureCommandClass []byte

	SecurityScheme0CommandClass []byte
}

func (cmd ApplicationNodeInfoReport) CommandClassID() cc.CommandClassID {
	return 0x5F
}

func (cmd ApplicationNodeInfoReport) CommandID() cc.CommandID {
	return CommandApplicationNodeInfoReport
}

func (cmd ApplicationNodeInfoReport) CommandIDString() string {
	return "COMMAND_APPLICATION_NODE_INFO_REPORT"
}

func (cmd *ApplicationNodeInfoReport) UnmarshalBinary(data []byte) error {
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

	{
		fieldStart := i
		for ; i < len(payload) && payload[i] != 0xF1; i++ {
		}
		cmd.NonSecureCommandClass = payload[fieldStart:i]
	}

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	i += 1 // skipping MARKER
	if len(payload) <= i {
		return nil
	}

	if len(payload) <= i {
		return nil
	}

	cmd.SecurityScheme0CommandClass = payload[i:]

	return nil
}

func (cmd *ApplicationNodeInfoReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		if cmd.NonSecureCommandClass != nil && len(cmd.NonSecureCommandClass) > 0 {
			payload = append(payload, cmd.NonSecureCommandClass...)
		}
		payload = append(payload, 0xF1)
	}

	payload = append(payload, 0xF1) // marker

	payload = append(payload, cmd.SecurityScheme0CommandClass...)

	return
}