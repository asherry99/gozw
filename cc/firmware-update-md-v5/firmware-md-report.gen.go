// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package firmwareupdatemdv5

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/gozwave/gozw/cc"
)

const CommandFirmwareMdReport cc.CommandID = 0x02

func init() {
	gob.Register(FirmwareMdReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x7A),
		Command:      cc.CommandID(0x02),
		Version:      5,
	}, NewFirmwareMdReport)
}

func NewFirmwareMdReport() cc.Command {
	return &FirmwareMdReport{}
}

// <no value>
type FirmwareMdReport struct {
	ManufacturerId uint16

	Firmware0Id uint16

	Firmware0Checksum uint16

	FirmwareUpgradable byte

	NumberOfFirmwareTargets byte

	MaxFragmentSize uint16

	HardwareVersion byte

	Vg1 []FirmwareMdReportVg1
}

type FirmwareMdReportVg1 struct {
	FirmwareId uint16
}

func (cmd FirmwareMdReport) CommandClassID() cc.CommandClassID {
	return 0x7A
}

func (cmd FirmwareMdReport) CommandID() cc.CommandID {
	return CommandFirmwareMdReport
}

func (cmd FirmwareMdReport) CommandIDString() string {
	return "FIRMWARE_MD_REPORT"
}

func (cmd *FirmwareMdReport) UnmarshalBinary(data []byte) error {
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

	cmd.ManufacturerId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Firmware0Id = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Firmware0Checksum = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.FirmwareUpgradable = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NumberOfFirmwareTargets = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MaxFragmentSize = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	for i < int(cmd.NumberOfFirmwareTargets) {

		vg1 := FirmwareMdReportVg1{}

		if len(payload) <= i {
			return errors.New("slice index out of bounds")
		}

		firmwareId := binary.BigEndian.Uint16(payload[i : i+2])
		i += 2

		vg1.FirmwareId = firmwareId

		cmd.Vg1 = append(cmd.Vg1, vg1)
	}

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.HardwareVersion = payload[i]
	i++

	return nil
}

func (cmd *FirmwareMdReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ManufacturerId)
		payload = append(payload, buf...)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.Firmware0Id)
		payload = append(payload, buf...)
	}

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.Firmware0Checksum)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.FirmwareUpgradable)

	payload = append(payload, cmd.NumberOfFirmwareTargets)

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.MaxFragmentSize)
		payload = append(payload, buf...)
	}

	for _, vg := range cmd.Vg1 {

		{
			buf := make([]byte, 2)
			binary.BigEndian.PutUint16(buf, vg.FirmwareId)
			payload = append(payload, buf...)
		}

	}

	payload = append(payload, cmd.HardwareVersion)

	return
}
