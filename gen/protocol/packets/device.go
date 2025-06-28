// Code generated. DO NOT EDIT.
// Source: https://github.com/LIFX/public-protocol@f061a60a2e9305e63ca75cfa53ad78d276de0e0a
// Generated: 2025-06-28T09:47:36Z
package packets

import (
	"bytes"
	"encoding/binary"

	"github.com/alessio-palumbo/lifxprotocol-go/gen/protocol/enums"
)

// DeviceGetService defines the fields for a packet of 0 bytes.
type DeviceGetService struct {
}

// PayloadType returns the LIFX message type for DeviceGetService.
func (p *DeviceGetService) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceGetService)
}

// Size is the total size of the message.
func (p *DeviceGetService) Size() int {
	return 0
}

// MarshalBinary encodes the DeviceGetService packet into LIFX binary format.
func (p *DeviceGetService) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceGetService struct.
func (p *DeviceGetService) UnmarshalBinary(data []byte) error {
	return nil
}

// DeviceStateService defines the fields for a packet of 5 bytes.
type DeviceStateService struct {
	Service enums.DeviceService
	Port    uint32
}

// PayloadType returns the LIFX message type for DeviceStateService.
func (p *DeviceStateService) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceStateService)
}

// Size is the total size of the message.
func (p *DeviceStateService) Size() int {
	return 5
}

// MarshalBinary encodes the DeviceStateService packet into LIFX binary format.
func (p *DeviceStateService) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Service); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Port); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceStateService struct.
func (p *DeviceStateService) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Service); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Port); err != nil {
		return err
	}
	return nil
}

// DeviceGetHostFirmware defines the fields for a packet of 0 bytes.
type DeviceGetHostFirmware struct {
}

// PayloadType returns the LIFX message type for DeviceGetHostFirmware.
func (p *DeviceGetHostFirmware) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceGetHostFirmware)
}

// Size is the total size of the message.
func (p *DeviceGetHostFirmware) Size() int {
	return 0
}

// MarshalBinary encodes the DeviceGetHostFirmware packet into LIFX binary format.
func (p *DeviceGetHostFirmware) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceGetHostFirmware struct.
func (p *DeviceGetHostFirmware) UnmarshalBinary(data []byte) error {
	return nil
}

// DeviceStateHostFirmware defines the fields for a packet of 20 bytes.
type DeviceStateHostFirmware struct {
	Build        uint64
	Reserved1    uint64
	VersionMinor uint16
	VersionMajor uint16
}

// PayloadType returns the LIFX message type for DeviceStateHostFirmware.
func (p *DeviceStateHostFirmware) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceStateHostFirmware)
}

// Size is the total size of the message.
func (p *DeviceStateHostFirmware) Size() int {
	return 20
}

// MarshalBinary encodes the DeviceStateHostFirmware packet into LIFX binary format.
func (p *DeviceStateHostFirmware) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Build); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved1); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.VersionMinor); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.VersionMajor); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceStateHostFirmware struct.
func (p *DeviceStateHostFirmware) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Build); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved1); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.VersionMinor); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.VersionMajor); err != nil {
		return err
	}
	return nil
}

// DeviceGetWifiInfo defines the fields for a packet of 0 bytes.
type DeviceGetWifiInfo struct {
}

// PayloadType returns the LIFX message type for DeviceGetWifiInfo.
func (p *DeviceGetWifiInfo) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceGetWifiInfo)
}

// Size is the total size of the message.
func (p *DeviceGetWifiInfo) Size() int {
	return 0
}

// MarshalBinary encodes the DeviceGetWifiInfo packet into LIFX binary format.
func (p *DeviceGetWifiInfo) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceGetWifiInfo struct.
func (p *DeviceGetWifiInfo) UnmarshalBinary(data []byte) error {
	return nil
}

// DeviceStateWifiInfo defines the fields for a packet of 14 bytes.
type DeviceStateWifiInfo struct {
	Signal    float32
	Reserved1 uint32
	Reserved2 uint32
	Reserved3 uint16
}

// PayloadType returns the LIFX message type for DeviceStateWifiInfo.
func (p *DeviceStateWifiInfo) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceStateWifiInfo)
}

// Size is the total size of the message.
func (p *DeviceStateWifiInfo) Size() int {
	return 14
}

// MarshalBinary encodes the DeviceStateWifiInfo packet into LIFX binary format.
func (p *DeviceStateWifiInfo) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Signal); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved1); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved2); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved3); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceStateWifiInfo struct.
func (p *DeviceStateWifiInfo) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Signal); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved1); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved2); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved3); err != nil {
		return err
	}
	return nil
}

// DeviceGetWifiFirmware defines the fields for a packet of 0 bytes.
type DeviceGetWifiFirmware struct {
}

// PayloadType returns the LIFX message type for DeviceGetWifiFirmware.
func (p *DeviceGetWifiFirmware) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceGetWifiFirmware)
}

// Size is the total size of the message.
func (p *DeviceGetWifiFirmware) Size() int {
	return 0
}

// MarshalBinary encodes the DeviceGetWifiFirmware packet into LIFX binary format.
func (p *DeviceGetWifiFirmware) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceGetWifiFirmware struct.
func (p *DeviceGetWifiFirmware) UnmarshalBinary(data []byte) error {
	return nil
}

// DeviceStateWifiFirmware defines the fields for a packet of 20 bytes.
type DeviceStateWifiFirmware struct {
	Build        uint64
	Reserved1    uint64
	VersionMinor uint16
	VersionMajor uint16
}

// PayloadType returns the LIFX message type for DeviceStateWifiFirmware.
func (p *DeviceStateWifiFirmware) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceStateWifiFirmware)
}

// Size is the total size of the message.
func (p *DeviceStateWifiFirmware) Size() int {
	return 20
}

// MarshalBinary encodes the DeviceStateWifiFirmware packet into LIFX binary format.
func (p *DeviceStateWifiFirmware) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Build); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved1); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.VersionMinor); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.VersionMajor); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceStateWifiFirmware struct.
func (p *DeviceStateWifiFirmware) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Build); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved1); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.VersionMinor); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.VersionMajor); err != nil {
		return err
	}
	return nil
}

// DeviceGetPower defines the fields for a packet of 0 bytes.
type DeviceGetPower struct {
}

// PayloadType returns the LIFX message type for DeviceGetPower.
func (p *DeviceGetPower) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceGetPower)
}

// Size is the total size of the message.
func (p *DeviceGetPower) Size() int {
	return 0
}

// MarshalBinary encodes the DeviceGetPower packet into LIFX binary format.
func (p *DeviceGetPower) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceGetPower struct.
func (p *DeviceGetPower) UnmarshalBinary(data []byte) error {
	return nil
}

// DeviceSetPower defines the fields for a packet of 2 bytes.
type DeviceSetPower struct {
	Level uint16
}

// PayloadType returns the LIFX message type for DeviceSetPower.
func (p *DeviceSetPower) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceSetPower)
}

// Size is the total size of the message.
func (p *DeviceSetPower) Size() int {
	return 2
}

// MarshalBinary encodes the DeviceSetPower packet into LIFX binary format.
func (p *DeviceSetPower) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Level); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceSetPower struct.
func (p *DeviceSetPower) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Level); err != nil {
		return err
	}
	return nil
}

// DeviceStatePower defines the fields for a packet of 2 bytes.
type DeviceStatePower struct {
	Level uint16
}

// PayloadType returns the LIFX message type for DeviceStatePower.
func (p *DeviceStatePower) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceStatePower)
}

// Size is the total size of the message.
func (p *DeviceStatePower) Size() int {
	return 2
}

// MarshalBinary encodes the DeviceStatePower packet into LIFX binary format.
func (p *DeviceStatePower) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Level); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceStatePower struct.
func (p *DeviceStatePower) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Level); err != nil {
		return err
	}
	return nil
}

// DeviceGetLabel defines the fields for a packet of 0 bytes.
type DeviceGetLabel struct {
}

// PayloadType returns the LIFX message type for DeviceGetLabel.
func (p *DeviceGetLabel) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceGetLabel)
}

// Size is the total size of the message.
func (p *DeviceGetLabel) Size() int {
	return 0
}

// MarshalBinary encodes the DeviceGetLabel packet into LIFX binary format.
func (p *DeviceGetLabel) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceGetLabel struct.
func (p *DeviceGetLabel) UnmarshalBinary(data []byte) error {
	return nil
}

// DeviceSetLabel defines the fields for a packet of 32 bytes.
type DeviceSetLabel struct {
	Label [32]byte
}

// PayloadType returns the LIFX message type for DeviceSetLabel.
func (p *DeviceSetLabel) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceSetLabel)
}

// Size is the total size of the message.
func (p *DeviceSetLabel) Size() int {
	return 32
}

// MarshalBinary encodes the DeviceSetLabel packet into LIFX binary format.
func (p *DeviceSetLabel) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Label); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceSetLabel struct.
func (p *DeviceSetLabel) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Label); err != nil {
		return err
	}
	return nil
}

// DeviceStateLabel defines the fields for a packet of 32 bytes.
type DeviceStateLabel struct {
	Label [32]byte
}

// PayloadType returns the LIFX message type for DeviceStateLabel.
func (p *DeviceStateLabel) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceStateLabel)
}

// Size is the total size of the message.
func (p *DeviceStateLabel) Size() int {
	return 32
}

// MarshalBinary encodes the DeviceStateLabel packet into LIFX binary format.
func (p *DeviceStateLabel) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Label); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceStateLabel struct.
func (p *DeviceStateLabel) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Label); err != nil {
		return err
	}
	return nil
}

// DeviceGetVersion defines the fields for a packet of 0 bytes.
type DeviceGetVersion struct {
}

// PayloadType returns the LIFX message type for DeviceGetVersion.
func (p *DeviceGetVersion) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceGetVersion)
}

// Size is the total size of the message.
func (p *DeviceGetVersion) Size() int {
	return 0
}

// MarshalBinary encodes the DeviceGetVersion packet into LIFX binary format.
func (p *DeviceGetVersion) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceGetVersion struct.
func (p *DeviceGetVersion) UnmarshalBinary(data []byte) error {
	return nil
}

// DeviceStateVersion defines the fields for a packet of 12 bytes.
type DeviceStateVersion struct {
	Vendor    uint32
	Product   uint32
	Reserved1 uint32
}

// PayloadType returns the LIFX message type for DeviceStateVersion.
func (p *DeviceStateVersion) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceStateVersion)
}

// Size is the total size of the message.
func (p *DeviceStateVersion) Size() int {
	return 12
}

// MarshalBinary encodes the DeviceStateVersion packet into LIFX binary format.
func (p *DeviceStateVersion) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Vendor); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Product); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved1); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceStateVersion struct.
func (p *DeviceStateVersion) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Vendor); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Product); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved1); err != nil {
		return err
	}
	return nil
}

// DeviceGetInfo defines the fields for a packet of 0 bytes.
type DeviceGetInfo struct {
}

// PayloadType returns the LIFX message type for DeviceGetInfo.
func (p *DeviceGetInfo) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceGetInfo)
}

// Size is the total size of the message.
func (p *DeviceGetInfo) Size() int {
	return 0
}

// MarshalBinary encodes the DeviceGetInfo packet into LIFX binary format.
func (p *DeviceGetInfo) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceGetInfo struct.
func (p *DeviceGetInfo) UnmarshalBinary(data []byte) error {
	return nil
}

// DeviceStateInfo defines the fields for a packet of 24 bytes.
type DeviceStateInfo struct {
	Time     uint64
	Uptime   uint64
	Downtime uint64
}

// PayloadType returns the LIFX message type for DeviceStateInfo.
func (p *DeviceStateInfo) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceStateInfo)
}

// Size is the total size of the message.
func (p *DeviceStateInfo) Size() int {
	return 24
}

// MarshalBinary encodes the DeviceStateInfo packet into LIFX binary format.
func (p *DeviceStateInfo) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Time); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Uptime); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Downtime); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceStateInfo struct.
func (p *DeviceStateInfo) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Time); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Uptime); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Downtime); err != nil {
		return err
	}
	return nil
}

// DeviceSetReboot defines the fields for a packet of 0 bytes.
type DeviceSetReboot struct {
}

// PayloadType returns the LIFX message type for DeviceSetReboot.
func (p *DeviceSetReboot) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceSetReboot)
}

// Size is the total size of the message.
func (p *DeviceSetReboot) Size() int {
	return 0
}

// MarshalBinary encodes the DeviceSetReboot packet into LIFX binary format.
func (p *DeviceSetReboot) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceSetReboot struct.
func (p *DeviceSetReboot) UnmarshalBinary(data []byte) error {
	return nil
}

// DeviceAcknowledgement defines the fields for a packet of 0 bytes.
type DeviceAcknowledgement struct {
}

// PayloadType returns the LIFX message type for DeviceAcknowledgement.
func (p *DeviceAcknowledgement) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceAcknowledgement)
}

// Size is the total size of the message.
func (p *DeviceAcknowledgement) Size() int {
	return 0
}

// MarshalBinary encodes the DeviceAcknowledgement packet into LIFX binary format.
func (p *DeviceAcknowledgement) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceAcknowledgement struct.
func (p *DeviceAcknowledgement) UnmarshalBinary(data []byte) error {
	return nil
}

// DeviceGetLocation defines the fields for a packet of 0 bytes.
type DeviceGetLocation struct {
}

// PayloadType returns the LIFX message type for DeviceGetLocation.
func (p *DeviceGetLocation) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceGetLocation)
}

// Size is the total size of the message.
func (p *DeviceGetLocation) Size() int {
	return 0
}

// MarshalBinary encodes the DeviceGetLocation packet into LIFX binary format.
func (p *DeviceGetLocation) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceGetLocation struct.
func (p *DeviceGetLocation) UnmarshalBinary(data []byte) error {
	return nil
}

// DeviceSetLocation defines the fields for a packet of 56 bytes.
type DeviceSetLocation struct {
	Location  [16]byte
	Label     [32]byte
	UpdatedAt uint64
}

// PayloadType returns the LIFX message type for DeviceSetLocation.
func (p *DeviceSetLocation) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceSetLocation)
}

// Size is the total size of the message.
func (p *DeviceSetLocation) Size() int {
	return 56
}

// MarshalBinary encodes the DeviceSetLocation packet into LIFX binary format.
func (p *DeviceSetLocation) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Location); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Label); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.UpdatedAt); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceSetLocation struct.
func (p *DeviceSetLocation) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Location); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Label); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.UpdatedAt); err != nil {
		return err
	}
	return nil
}

// DeviceStateLocation defines the fields for a packet of 56 bytes.
type DeviceStateLocation struct {
	Location  [16]byte
	Label     [32]byte
	UpdatedAt uint64
}

// PayloadType returns the LIFX message type for DeviceStateLocation.
func (p *DeviceStateLocation) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceStateLocation)
}

// Size is the total size of the message.
func (p *DeviceStateLocation) Size() int {
	return 56
}

// MarshalBinary encodes the DeviceStateLocation packet into LIFX binary format.
func (p *DeviceStateLocation) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Location); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Label); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.UpdatedAt); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceStateLocation struct.
func (p *DeviceStateLocation) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Location); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Label); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.UpdatedAt); err != nil {
		return err
	}
	return nil
}

// DeviceGetGroup defines the fields for a packet of 0 bytes.
type DeviceGetGroup struct {
}

// PayloadType returns the LIFX message type for DeviceGetGroup.
func (p *DeviceGetGroup) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceGetGroup)
}

// Size is the total size of the message.
func (p *DeviceGetGroup) Size() int {
	return 0
}

// MarshalBinary encodes the DeviceGetGroup packet into LIFX binary format.
func (p *DeviceGetGroup) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceGetGroup struct.
func (p *DeviceGetGroup) UnmarshalBinary(data []byte) error {
	return nil
}

// DeviceSetGroup defines the fields for a packet of 56 bytes.
type DeviceSetGroup struct {
	Group     [16]byte
	Label     [32]byte
	UpdatedAt uint64
}

// PayloadType returns the LIFX message type for DeviceSetGroup.
func (p *DeviceSetGroup) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceSetGroup)
}

// Size is the total size of the message.
func (p *DeviceSetGroup) Size() int {
	return 56
}

// MarshalBinary encodes the DeviceSetGroup packet into LIFX binary format.
func (p *DeviceSetGroup) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Group); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Label); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.UpdatedAt); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceSetGroup struct.
func (p *DeviceSetGroup) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Group); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Label); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.UpdatedAt); err != nil {
		return err
	}
	return nil
}

// DeviceStateGroup defines the fields for a packet of 56 bytes.
type DeviceStateGroup struct {
	Group     [16]byte
	Label     [32]byte
	UpdatedAt uint64
}

// PayloadType returns the LIFX message type for DeviceStateGroup.
func (p *DeviceStateGroup) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceStateGroup)
}

// Size is the total size of the message.
func (p *DeviceStateGroup) Size() int {
	return 56
}

// MarshalBinary encodes the DeviceStateGroup packet into LIFX binary format.
func (p *DeviceStateGroup) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Group); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Label); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.UpdatedAt); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceStateGroup struct.
func (p *DeviceStateGroup) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Group); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Label); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.UpdatedAt); err != nil {
		return err
	}
	return nil
}

// DeviceEchoRequest defines the fields for a packet of 64 bytes.
type DeviceEchoRequest struct {
	Payload [64]byte
}

// PayloadType returns the LIFX message type for DeviceEchoRequest.
func (p *DeviceEchoRequest) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceEchoRequest)
}

// Size is the total size of the message.
func (p *DeviceEchoRequest) Size() int {
	return 64
}

// MarshalBinary encodes the DeviceEchoRequest packet into LIFX binary format.
func (p *DeviceEchoRequest) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Payload); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceEchoRequest struct.
func (p *DeviceEchoRequest) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Payload); err != nil {
		return err
	}
	return nil
}

// DeviceEchoResponse defines the fields for a packet of 64 bytes.
type DeviceEchoResponse struct {
	Payload [64]byte
}

// PayloadType returns the LIFX message type for DeviceEchoResponse.
func (p *DeviceEchoResponse) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceEchoResponse)
}

// Size is the total size of the message.
func (p *DeviceEchoResponse) Size() int {
	return 64
}

// MarshalBinary encodes the DeviceEchoResponse packet into LIFX binary format.
func (p *DeviceEchoResponse) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Payload); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceEchoResponse struct.
func (p *DeviceEchoResponse) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Payload); err != nil {
		return err
	}
	return nil
}

// DeviceStateUnhandled defines the fields for a packet of 2 bytes.
type DeviceStateUnhandled struct {
	UnhandledType uint16
}

// PayloadType returns the LIFX message type for DeviceStateUnhandled.
func (p *DeviceStateUnhandled) PayloadType() uint16 {
	return uint16(PayloadTypeDeviceStateUnhandled)
}

// Size is the total size of the message.
func (p *DeviceStateUnhandled) Size() int {
	return 2
}

// MarshalBinary encodes the DeviceStateUnhandled packet into LIFX binary format.
func (p *DeviceStateUnhandled) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.UnhandledType); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the DeviceStateUnhandled struct.
func (p *DeviceStateUnhandled) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.UnhandledType); err != nil {
		return err
	}
	return nil
}
