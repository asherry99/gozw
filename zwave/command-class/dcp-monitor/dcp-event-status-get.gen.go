// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package dcpmonitor

import "encoding/binary"

// <no value>

type DcpEventStatusGet struct {
	Year uint16

	Month byte

	Day byte

	HourLocalTime byte

	MinuteLocalTime byte

	SecondLocalTime byte
}

func ParseDcpEventStatusGet(payload []byte) DcpEventStatusGet {
	val := DcpEventStatusGet{}

	i := 2

	val.Year = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	val.Month = payload[i]
	i++

	val.Day = payload[i]
	i++

	val.HourLocalTime = payload[i]
	i++

	val.MinuteLocalTime = payload[i]
	i++

	val.SecondLocalTime = payload[i]
	i++

	return val
}