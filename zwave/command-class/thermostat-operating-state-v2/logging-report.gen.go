// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatoperatingstatev2

// <no value>

type ThermostatOperatingStateLoggingReport struct {
	ReportsToFollow byte
}

func ParseThermostatOperatingStateLoggingReport(payload []byte) ThermostatOperatingStateLoggingReport {
	val := ThermostatOperatingStateLoggingReport{}

	i := 2

	val.ReportsToFollow = payload[i]
	i++

	return val
}