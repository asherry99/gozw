// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package avcontentdirectorymd

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandAvContentBrowseMdByLetterReport cc.CommandID = 0x04

func init() {
	gob.Register(AvContentBrowseMdByLetterReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x95),
		Command:      cc.CommandID(0x04),
		Version:      1,
	}, NewAvContentBrowseMdByLetterReport)
}

func NewAvContentBrowseMdByLetterReport() cc.Command {
	return &AvContentBrowseMdByLetterReport{}
}

// <no value>
type AvContentBrowseMdByLetterReport struct {
}

func (cmd AvContentBrowseMdByLetterReport) CommandClassID() cc.CommandClassID {
	return 0x95
}

func (cmd AvContentBrowseMdByLetterReport) CommandID() cc.CommandID {
	return CommandAvContentBrowseMdByLetterReport
}

func (cmd AvContentBrowseMdByLetterReport) CommandIDString() string {
	return "AV_CONTENT_BROWSE_MD_BY_LETTER_REPORT"
}

func (cmd *AvContentBrowseMdByLetterReport) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *AvContentBrowseMdByLetterReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}