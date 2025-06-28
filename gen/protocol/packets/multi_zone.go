// Code generated. DO NOT EDIT.
// Source: https://github.com/LIFX/public-protocol@f061a60a2e9305e63ca75cfa53ad78d276de0e0a
// Generated: 2025-06-28T09:47:36Z
package packets

import (
	"bytes"
	"encoding/binary"

	"github.com/alessio-palumbo/lifxprotocol-go/gen/protocol/enums"
)

// MultiZoneSetColorZones defines the fields for a packet of 15 bytes.
type MultiZoneSetColorZones struct {
	StartIndex uint8
	EndIndex   uint8
	Color      LightHsbk
	Duration   uint32
	Apply      enums.MultiZoneApplicationRequest
}

// PayloadType returns the LIFX message type for MultiZoneSetColorZones.
func (p *MultiZoneSetColorZones) PayloadType() uint16 {
	return uint16(PayloadTypeMultiZoneSetColorZones)
}

// Size is the total size of the message.
func (p *MultiZoneSetColorZones) Size() int {
	return 15
}

// MarshalBinary encodes the MultiZoneSetColorZones packet into LIFX binary format.
func (p *MultiZoneSetColorZones) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.StartIndex); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.EndIndex); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Color); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Duration); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Apply); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the MultiZoneSetColorZones struct.
func (p *MultiZoneSetColorZones) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.StartIndex); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.EndIndex); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Color); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Duration); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Apply); err != nil {
		return err
	}
	return nil
}

// MultiZoneGetColorZones defines the fields for a packet of 2 bytes.
type MultiZoneGetColorZones struct {
	StartIndex uint8
	EndIndex   uint8
}

// PayloadType returns the LIFX message type for MultiZoneGetColorZones.
func (p *MultiZoneGetColorZones) PayloadType() uint16 {
	return uint16(PayloadTypeMultiZoneGetColorZones)
}

// Size is the total size of the message.
func (p *MultiZoneGetColorZones) Size() int {
	return 2
}

// MarshalBinary encodes the MultiZoneGetColorZones packet into LIFX binary format.
func (p *MultiZoneGetColorZones) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.StartIndex); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.EndIndex); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the MultiZoneGetColorZones struct.
func (p *MultiZoneGetColorZones) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.StartIndex); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.EndIndex); err != nil {
		return err
	}
	return nil
}

// MultiZoneStateZone defines the fields for a packet of 10 bytes.
type MultiZoneStateZone struct {
	Count uint8
	Index uint8
	Color LightHsbk
}

// PayloadType returns the LIFX message type for MultiZoneStateZone.
func (p *MultiZoneStateZone) PayloadType() uint16 {
	return uint16(PayloadTypeMultiZoneStateZone)
}

// Size is the total size of the message.
func (p *MultiZoneStateZone) Size() int {
	return 10
}

// MarshalBinary encodes the MultiZoneStateZone packet into LIFX binary format.
func (p *MultiZoneStateZone) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Count); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Index); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Color); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the MultiZoneStateZone struct.
func (p *MultiZoneStateZone) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Count); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Index); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Color); err != nil {
		return err
	}
	return nil
}

// MultiZoneStateMultiZone defines the fields for a packet of 66 bytes.
type MultiZoneStateMultiZone struct {
	Count  uint8
	Index  uint8
	Colors [8]LightHsbk
}

// PayloadType returns the LIFX message type for MultiZoneStateMultiZone.
func (p *MultiZoneStateMultiZone) PayloadType() uint16 {
	return uint16(PayloadTypeMultiZoneStateMultiZone)
}

// Size is the total size of the message.
func (p *MultiZoneStateMultiZone) Size() int {
	return 66
}

// MarshalBinary encodes the MultiZoneStateMultiZone packet into LIFX binary format.
func (p *MultiZoneStateMultiZone) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Count); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Index); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Colors); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the MultiZoneStateMultiZone struct.
func (p *MultiZoneStateMultiZone) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Count); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Index); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Colors); err != nil {
		return err
	}
	return nil
}

// MultiZoneGetEffect defines the fields for a packet of 0 bytes.
type MultiZoneGetEffect struct {
}

// PayloadType returns the LIFX message type for MultiZoneGetEffect.
func (p *MultiZoneGetEffect) PayloadType() uint16 {
	return uint16(PayloadTypeMultiZoneGetEffect)
}

// Size is the total size of the message.
func (p *MultiZoneGetEffect) Size() int {
	return 0
}

// MarshalBinary encodes the MultiZoneGetEffect packet into LIFX binary format.
func (p *MultiZoneGetEffect) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the MultiZoneGetEffect struct.
func (p *MultiZoneGetEffect) UnmarshalBinary(data []byte) error {
	return nil
}

// MultiZoneSetEffect defines the fields for a packet of 59 bytes.
type MultiZoneSetEffect struct {
	Settings MultiZoneEffectSettings
}

// PayloadType returns the LIFX message type for MultiZoneSetEffect.
func (p *MultiZoneSetEffect) PayloadType() uint16 {
	return uint16(PayloadTypeMultiZoneSetEffect)
}

// Size is the total size of the message.
func (p *MultiZoneSetEffect) Size() int {
	return 59
}

// MarshalBinary encodes the MultiZoneSetEffect packet into LIFX binary format.
func (p *MultiZoneSetEffect) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Settings); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the MultiZoneSetEffect struct.
func (p *MultiZoneSetEffect) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Settings); err != nil {
		return err
	}
	return nil
}

// MultiZoneStateEffect defines the fields for a packet of 59 bytes.
type MultiZoneStateEffect struct {
	Settings MultiZoneEffectSettings
}

// PayloadType returns the LIFX message type for MultiZoneStateEffect.
func (p *MultiZoneStateEffect) PayloadType() uint16 {
	return uint16(PayloadTypeMultiZoneStateEffect)
}

// Size is the total size of the message.
func (p *MultiZoneStateEffect) Size() int {
	return 59
}

// MarshalBinary encodes the MultiZoneStateEffect packet into LIFX binary format.
func (p *MultiZoneStateEffect) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Settings); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the MultiZoneStateEffect struct.
func (p *MultiZoneStateEffect) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Settings); err != nil {
		return err
	}
	return nil
}

// MultiZoneExtendedSetColorZones defines the fields for a packet of 664 bytes.
type MultiZoneExtendedSetColorZones struct {
	Duration    uint32
	Apply       enums.MultiZoneExtendedApplicationRequest
	Index       uint16
	ColorsCount uint8
	Colors      [82]LightHsbk
}

// PayloadType returns the LIFX message type for MultiZoneExtendedSetColorZones.
func (p *MultiZoneExtendedSetColorZones) PayloadType() uint16 {
	return uint16(PayloadTypeMultiZoneExtendedSetColorZones)
}

// Size is the total size of the message.
func (p *MultiZoneExtendedSetColorZones) Size() int {
	return 664
}

// MarshalBinary encodes the MultiZoneExtendedSetColorZones packet into LIFX binary format.
func (p *MultiZoneExtendedSetColorZones) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Duration); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Apply); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Index); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.ColorsCount); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Colors); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the MultiZoneExtendedSetColorZones struct.
func (p *MultiZoneExtendedSetColorZones) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Duration); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Apply); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Index); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.ColorsCount); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Colors); err != nil {
		return err
	}
	return nil
}

// MultiZoneExtendedGetColorZones defines the fields for a packet of 0 bytes.
type MultiZoneExtendedGetColorZones struct {
}

// PayloadType returns the LIFX message type for MultiZoneExtendedGetColorZones.
func (p *MultiZoneExtendedGetColorZones) PayloadType() uint16 {
	return uint16(PayloadTypeMultiZoneExtendedGetColorZones)
}

// Size is the total size of the message.
func (p *MultiZoneExtendedGetColorZones) Size() int {
	return 0
}

// MarshalBinary encodes the MultiZoneExtendedGetColorZones packet into LIFX binary format.
func (p *MultiZoneExtendedGetColorZones) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the MultiZoneExtendedGetColorZones struct.
func (p *MultiZoneExtendedGetColorZones) UnmarshalBinary(data []byte) error {
	return nil
}

// MultiZoneExtendedStateMultiZone defines the fields for a packet of 661 bytes.
type MultiZoneExtendedStateMultiZone struct {
	Count       uint16
	Index       uint16
	ColorsCount uint8
	Colors      [82]LightHsbk
}

// PayloadType returns the LIFX message type for MultiZoneExtendedStateMultiZone.
func (p *MultiZoneExtendedStateMultiZone) PayloadType() uint16 {
	return uint16(PayloadTypeMultiZoneExtendedStateMultiZone)
}

// Size is the total size of the message.
func (p *MultiZoneExtendedStateMultiZone) Size() int {
	return 661
}

// MarshalBinary encodes the MultiZoneExtendedStateMultiZone packet into LIFX binary format.
func (p *MultiZoneExtendedStateMultiZone) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Count); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Index); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.ColorsCount); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Colors); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the MultiZoneExtendedStateMultiZone struct.
func (p *MultiZoneExtendedStateMultiZone) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Count); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Index); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.ColorsCount); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Colors); err != nil {
		return err
	}
	return nil
}
