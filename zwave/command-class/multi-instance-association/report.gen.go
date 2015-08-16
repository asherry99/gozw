// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package multiinstanceassociation

import "errors"

// <no value>

type MultiInstanceAssociationReport struct {
	GroupingIdentifier byte

	MaxNodesSupported byte

	ReportsToFollow byte

	NodeId []byte
}

func (cmd *MultiInstanceAssociationReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GroupingIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MaxNodesSupported = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ReportsToFollow = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	{
		markerIndex := i
		for ; markerIndex < len(payload) && payload[markerIndex] != 0x00; markerIndex++ {
		}
		val.NodeId = payload[i:markerIndex]
	}

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	i += 1 // skipping MARKER

	return nil
}