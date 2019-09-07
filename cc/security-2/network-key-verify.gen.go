// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package security2

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandNetworkKeyVerify cc.CommandID = 0x0B

func init() {
	gob.Register(NetworkKeyVerify{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x9F),
		Command:      cc.CommandID(0x0B),
		Version:      1,
	}, NewNetworkKeyVerify)
}

func NewNetworkKeyVerify() cc.Command {
	return &NetworkKeyVerify{}
}

// <no value>
type NetworkKeyVerify struct {
}

func (cmd NetworkKeyVerify) CommandClassID() cc.CommandClassID {
	return 0x9F
}

func (cmd NetworkKeyVerify) CommandID() cc.CommandID {
	return CommandNetworkKeyVerify
}

func (cmd NetworkKeyVerify) CommandIDString() string {
	return "SECURITY_2_NETWORK_KEY_VERIFY"
}

func (cmd *NetworkKeyVerify) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *NetworkKeyVerify) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
