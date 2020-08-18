package serialapi

import (
	"fmt"

	"github.com/gozwave/gozw/frame"
	"github.com/gozwave/gozw/protocol"
	"github.com/gozwave/gozw/session"
)

// RequestNodeNeighborUpdate will requests nodes update there nearest neighbor tables.
func (s *Layer) RequestNodeNeighborUpdate(nodeID byte) error {

	done := make(chan *frame.Frame)

	request := &session.Request{
		FunctionID: protocol.FnRequestNodeNeighborUpdate,
		Payload:    []byte{nodeID},
		HasReturn:  true,
		ReturnCallback: func(err error, ret *frame.Frame) bool {
			done <- ret
			return false
		},
	}

	s.sessionLayer.MakeRequest(request)
	ret := <-done

	if ret == nil {
		return fmt.Errorf("Error updating nearest neighbor mappings for node %d", nodeID)
	}

	return nil
}
