// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package simpleavcontrol

// <no value>

type SimpleAvControlSupportedGet struct {
	ReportNo byte
}

func ParseSimpleAvControlSupportedGet(payload []byte) SimpleAvControlSupportedGet {
	val := SimpleAvControlSupportedGet{}

	i := 2

	val.ReportNo = payload[i]
	i++

	return val
}