package serialapi

import (
	"errors"
	"fmt"
	"time"

	"github.com/bjyoungblood/gozw/zwave/frame"
	"github.com/bjyoungblood/gozw/zwave/protocol"
	"github.com/bjyoungblood/gozw/zwave/session"
	"github.com/davecgh/go-spew/spew"
)

func (s *SerialAPILayer) AddNode() (*AddRemoveNodeCallback, error) {

	var newNode *AddRemoveNodeCallback

	addNodeDone := make(chan bool)
	done := make(chan *frame.Frame)

	request := &session.Request{
		FunctionId: protocol.FnAddNodeToNetwork,
		Payload:    []byte{protocol.AddNodeAny | protocol.AddNodeOptionNetworkWide | protocol.AddNodeOptionNormalPower},

		HasReturn:        false,
		ReceivesCallback: true,
		Lock:             true,
		Timeout:          60 * time.Second,
		Release:          addNodeDone,

		Callback: func(cbFrame frame.Frame) {
			cbData := parseAddRemoveNodeCallback(cbFrame.Payload)

			switch cbData.Status {
			case protocol.AddNodeStatusLearnReady:
				fmt.Println("ADD NODE: learn ready")

			case protocol.AddNodeStatusNodeFound:
				fmt.Println("ADD NODE: node found")

			case protocol.AddNodeStatusAddingSlave:
				fmt.Println("ADD NODE: adding slave node")
				newNode = cbData

			case protocol.AddNodeStatusAddingController:
				// hey, i just met you, and this is crazy
				// but it could happen, so implement me maybe
				fmt.Println("ADD NODE: adding controller node")
				newNode = cbData

			case protocol.AddNodeStatusProtocolDone:
				fmt.Println("ADD NODE: protocol done")
				reply := addRemoveStatusFrame(
					protocol.FnAddNodeToNetwork,
					protocol.AddNodeStop,
					cbData.CallbackId,
				)
				s.sessionLayer.SendFrameDirect(reply)

			case protocol.AddNodeStatusDone:
				fmt.Println("ADD NODE: done")
				spew.Dump(cbData)
				reply := addRemoveStatusFrame(
					protocol.FnAddNodeToNetwork,
					protocol.AddNodeStop,
					0,
				)
				s.sessionLayer.SendFrameDirect(reply)

				addNodeDone <- true
				close(addNodeDone)
				done <- &cbFrame

			case protocol.AddNodeStatusFailed:
				fmt.Println("ADD NODE: failed")
				reply := addRemoveStatusFrame(
					protocol.FnAddNodeToNetwork,
					protocol.AddNodeStop,
					cbData.CallbackId,
				)
				s.sessionLayer.SendFrameDirect(reply)

			default:
				fmt.Println("ADD NODE: unknown status", cbData.Status)
			}
		},
	}

	s.sessionLayer.MakeRequest(request)
	ret := <-done

	if ret == nil {
		return nil, errors.New("Error adding node")
	}

	return newNode, nil

}

func (s *SerialAPILayer) RemoveNode() (*AddRemoveNodeCallback, error) {

	var removedNode *AddRemoveNodeCallback

	removeNodeDone := make(chan bool)
	done := make(chan *frame.Frame)

	request := &session.Request{
		FunctionId: protocol.FnRemoveNodeFromNetwork,
		Payload:    []byte{protocol.RemoveNodeAny | protocol.RemoveNodeOptionNetworkWide | protocol.RemoveNodeOptionNormalPower},

		HasReturn:        false,
		ReceivesCallback: true,
		Lock:             true,
		Timeout:          60 * time.Second,
		Release:          removeNodeDone,

		Callback: func(cbFrame frame.Frame) {
			cbData := parseAddRemoveNodeCallback(cbFrame.Payload)

			switch cbData.Status {
			case protocol.RemoveNodeStatusLearnReady:
				fmt.Println("REMOVE NODE: learn ready")

			case protocol.RemoveNodeStatusNodeFound:
				fmt.Println("REMOVE NODE: node found")

			case protocol.RemoveNodeStatusRemovingSlave:
				fmt.Println("REMOVE NODE: removing slave node")
				removedNode = cbData

			case protocol.RemoveNodeStatusRemovingController:
				// hey, i just met you, and this is crazy
				// but it could happen, so implement me maybe
				fmt.Println("REMOVE NODE: removing controller node")
				removedNode = cbData

			case protocol.RemoveNodeStatusProtocolDone:
				fmt.Println("REMOVE NODE: protocol done")
				reply := addRemoveStatusFrame(
					protocol.FnRemoveNodeFromNetwork,
					protocol.RemoveNodeStop,
					cbData.CallbackId,
				)
				s.sessionLayer.SendFrameDirect(reply)

			case protocol.RemoveNodeStatusDone:
				fmt.Println("REMOVE NODE: done")
				reply := addRemoveStatusFrame(
					protocol.FnRemoveNodeFromNetwork,
					protocol.RemoveNodeStop,
					0,
				)
				s.sessionLayer.SendFrameDirect(reply)

				removeNodeDone <- true
				close(removeNodeDone)
				done <- &cbFrame

			case protocol.RemoveNodeStatusFailed:
				fmt.Println("REMOVE NODE: failed")
				reply := addRemoveStatusFrame(
					protocol.FnRemoveNodeFromNetwork,
					protocol.RemoveNodeStop,
					cbData.CallbackId,
				)
				s.sessionLayer.SendFrameDirect(reply)

			default:
				fmt.Println("REMOVE NODE: unknown status", cbData.Status)
			}
		},
	}

	s.sessionLayer.MakeRequest(request)
	ret := <-done

	if ret == nil {
		return nil, errors.New("Error removing node")
	}

	return removedNode, nil

}

func addRemoveStatusFrame(functionId, status, callbackId uint8) *frame.Frame {
	return frame.NewRequestFrame([]byte{
		functionId,
		status,
		callbackId,
	})
}

type AddRemoveNodeCallback struct {
	CommandId      byte
	CallbackId     byte
	Status         byte
	Source         byte
	Length         byte
	Basic          uint8
	Generic        uint8
	Specific       uint8
	CommandClasses []byte
}

func parseAddRemoveNodeCallback(payload []byte) *AddRemoveNodeCallback {
	val := &AddRemoveNodeCallback{
		CommandId:  payload[0],
		CallbackId: payload[1],
		Status:     payload[2],
		Source:     payload[3],
		Length:     payload[4],
	}

	if val.Length == 0 {
		return val
	}

	if val.Length >= 1 {
		val.Basic = payload[5]
	}

	if val.Length >= 2 {
		val.Generic = payload[6]
	}

	if val.Length >= 3 {
		val.Specific = payload[7]
	}

	if val.Length >= 4 {
		val.CommandClasses = payload[8:]
	}

	return val
}
