// Code generated. DO NOT EDIT.
// Source: https://github.com/LIFX/public-protocol@f061a60a2e9305e63ca75cfa53ad78d276de0e0a
// Generated: 2025-06-28T09:47:36Z
package packets

import (
	"bytes"
	"encoding/binary"
)

// TileGetDeviceChain defines the fields for a packet of 0 bytes.
type TileGetDeviceChain struct {
}

// PayloadType returns the LIFX message type for TileGetDeviceChain.
func (p *TileGetDeviceChain) PayloadType() uint16 {
	return uint16(PayloadTypeTileGetDeviceChain)
}

// Size is the total size of the message.
func (p *TileGetDeviceChain) Size() int {
	return 0
}

// MarshalBinary encodes the TileGetDeviceChain packet into LIFX binary format.
func (p *TileGetDeviceChain) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the TileGetDeviceChain struct.
func (p *TileGetDeviceChain) UnmarshalBinary(data []byte) error {
	return nil
}

// TileStateDeviceChain defines the fields for a packet of 882 bytes.
type TileStateDeviceChain struct {
	StartIndex       uint8
	TileDevices      [16]TileStateDevice
	TileDevicesCount uint8
}

// PayloadType returns the LIFX message type for TileStateDeviceChain.
func (p *TileStateDeviceChain) PayloadType() uint16 {
	return uint16(PayloadTypeTileStateDeviceChain)
}

// Size is the total size of the message.
func (p *TileStateDeviceChain) Size() int {
	return 882
}

// MarshalBinary encodes the TileStateDeviceChain packet into LIFX binary format.
func (p *TileStateDeviceChain) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.StartIndex); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.TileDevices); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.TileDevicesCount); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the TileStateDeviceChain struct.
func (p *TileStateDeviceChain) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.StartIndex); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.TileDevices); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.TileDevicesCount); err != nil {
		return err
	}
	return nil
}

// TileSetUserPosition defines the fields for a packet of 11 bytes.
type TileSetUserPosition struct {
	TileIndex uint8
	Reserved1 uint8
	Reserved2 uint8
	UserX     float32
	UserY     float32
}

// PayloadType returns the LIFX message type for TileSetUserPosition.
func (p *TileSetUserPosition) PayloadType() uint16 {
	return uint16(PayloadTypeTileSetUserPosition)
}

// Size is the total size of the message.
func (p *TileSetUserPosition) Size() int {
	return 11
}

// MarshalBinary encodes the TileSetUserPosition packet into LIFX binary format.
func (p *TileSetUserPosition) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.TileIndex); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved1); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved2); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.UserX); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.UserY); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the TileSetUserPosition struct.
func (p *TileSetUserPosition) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.TileIndex); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved1); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved2); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.UserX); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.UserY); err != nil {
		return err
	}
	return nil
}

// TileGet64 defines the fields for a packet of 6 bytes.
type TileGet64 struct {
	TileIndex uint8
	Length    uint8
	Rect      TileBufferRect
}

// PayloadType returns the LIFX message type for TileGet64.
func (p *TileGet64) PayloadType() uint16 {
	return uint16(PayloadTypeTileGet64)
}

// Size is the total size of the message.
func (p *TileGet64) Size() int {
	return 6
}

// MarshalBinary encodes the TileGet64 packet into LIFX binary format.
func (p *TileGet64) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.TileIndex); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Length); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Rect); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the TileGet64 struct.
func (p *TileGet64) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.TileIndex); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Length); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Rect); err != nil {
		return err
	}
	return nil
}

// TileState64 defines the fields for a packet of 517 bytes.
type TileState64 struct {
	TileIndex uint8
	Rect      TileBufferRect
	Colors    [64]LightHsbk
}

// PayloadType returns the LIFX message type for TileState64.
func (p *TileState64) PayloadType() uint16 {
	return uint16(PayloadTypeTileState64)
}

// Size is the total size of the message.
func (p *TileState64) Size() int {
	return 517
}

// MarshalBinary encodes the TileState64 packet into LIFX binary format.
func (p *TileState64) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.TileIndex); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Rect); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Colors); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the TileState64 struct.
func (p *TileState64) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.TileIndex); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Rect); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Colors); err != nil {
		return err
	}
	return nil
}

// TileSet64 defines the fields for a packet of 522 bytes.
type TileSet64 struct {
	TileIndex uint8
	Length    uint8
	Rect      TileBufferRect
	Duration  uint32
	Colors    [64]LightHsbk
}

// PayloadType returns the LIFX message type for TileSet64.
func (p *TileSet64) PayloadType() uint16 {
	return uint16(PayloadTypeTileSet64)
}

// Size is the total size of the message.
func (p *TileSet64) Size() int {
	return 522
}

// MarshalBinary encodes the TileSet64 packet into LIFX binary format.
func (p *TileSet64) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.TileIndex); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Length); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Rect); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Duration); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Colors); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the TileSet64 struct.
func (p *TileSet64) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.TileIndex); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Length); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Rect); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Duration); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Colors); err != nil {
		return err
	}
	return nil
}

// TileGetEffect defines the fields for a packet of 2 bytes.
type TileGetEffect struct {
	Reserved1 uint8
	Reserved2 uint8
}

// PayloadType returns the LIFX message type for TileGetEffect.
func (p *TileGetEffect) PayloadType() uint16 {
	return uint16(PayloadTypeTileGetEffect)
}

// Size is the total size of the message.
func (p *TileGetEffect) Size() int {
	return 2
}

// MarshalBinary encodes the TileGetEffect packet into LIFX binary format.
func (p *TileGetEffect) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved1); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved2); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the TileGetEffect struct.
func (p *TileGetEffect) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved1); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved2); err != nil {
		return err
	}
	return nil
}

// TileSetEffect defines the fields for a packet of 188 bytes.
type TileSetEffect struct {
	Reserved1 uint8
	Reserved2 uint8
	Settings  TileEffectSettings
}

// PayloadType returns the LIFX message type for TileSetEffect.
func (p *TileSetEffect) PayloadType() uint16 {
	return uint16(PayloadTypeTileSetEffect)
}

// Size is the total size of the message.
func (p *TileSetEffect) Size() int {
	return 188
}

// MarshalBinary encodes the TileSetEffect packet into LIFX binary format.
func (p *TileSetEffect) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved1); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved2); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Settings); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the TileSetEffect struct.
func (p *TileSetEffect) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved1); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved2); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Settings); err != nil {
		return err
	}
	return nil
}

// TileStateEffect defines the fields for a packet of 187 bytes.
type TileStateEffect struct {
	Reserved1 uint8
	Settings  TileEffectSettings
}

// PayloadType returns the LIFX message type for TileStateEffect.
func (p *TileStateEffect) PayloadType() uint16 {
	return uint16(PayloadTypeTileStateEffect)
}

// Size is the total size of the message.
func (p *TileStateEffect) Size() int {
	return 187
}

// MarshalBinary encodes the TileStateEffect packet into LIFX binary format.
func (p *TileStateEffect) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved1); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Settings); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the TileStateEffect struct.
func (p *TileStateEffect) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved1); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Settings); err != nil {
		return err
	}
	return nil
}
