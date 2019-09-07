// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package multichannelv2

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandEndPointGet cc.CommandID = 0x07

func init() {
	gob.Register(EndPointGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x60),
		Command:      cc.CommandID(0x07),
		Version:      2,
	}, NewEndPointGet)
}

func NewEndPointGet() cc.Command {
	return &EndPointGet{}
}

// <no value>
type EndPointGet struct {
}

func (cmd EndPointGet) CommandClassID() cc.CommandClassID {
	return 0x60
}

func (cmd EndPointGet) CommandID() cc.CommandID {
	return CommandEndPointGet
}

func (cmd EndPointGet) CommandIDString() string {
	return "MULTI_CHANNEL_END_POINT_GET"
}

func (cmd *EndPointGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *EndPointGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
