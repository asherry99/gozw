// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package networkmanagementproxyv2

import (
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandNodeInfoCachedGet cc.CommandID = 0x03

func init() {
	gob.Register(NodeInfoCachedGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x52),
		Command:      cc.CommandID(0x03),
		Version:      2,
	}, NewNodeInfoCachedGet)
}

func NewNodeInfoCachedGet() cc.Command {
	return &NodeInfoCachedGet{}
}

// <no value>
type NodeInfoCachedGet struct {
	SeqNo byte

	Properties1 struct {
		MaxAge byte
	}

	NodeId byte
}

func (cmd NodeInfoCachedGet) CommandClassID() cc.CommandClassID {
	return 0x52
}

func (cmd NodeInfoCachedGet) CommandID() cc.CommandID {
	return CommandNodeInfoCachedGet
}

func (cmd NodeInfoCachedGet) CommandIDString() string {
	return "NODE_INFO_CACHED_GET"
}

func (cmd *NodeInfoCachedGet) UnmarshalBinary(data []byte) error {
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

	cmd.Properties1.MaxAge = (payload[i] & 0x0F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NodeId = payload[i]
	i++

	return nil
}

func (cmd *NodeInfoCachedGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SeqNo)

	{
		var val byte

		val |= (cmd.Properties1.MaxAge) & byte(0x0F)

		payload = append(payload, val)
	}

	payload = append(payload, cmd.NodeId)

	return
}

func Noop() {
	// does nothing, just here to allow
	// consumers to invoke the init function
}