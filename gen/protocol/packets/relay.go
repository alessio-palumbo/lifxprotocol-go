// Code generated. DO NOT EDIT.
// Source: https://github.com/LIFX/public-protocol@f061a60a2e9305e63ca75cfa53ad78d276de0e0a
// Generated: 2025-06-28T09:47:36Z
package packets

import (
	"bytes"
	"encoding/binary"
)

// RelayGetPower defines the fields for a packet of 1 bytes.
type RelayGetPower struct {
	RelayIndex uint8
}

// PayloadType returns the LIFX message type for RelayGetPower.
func (p *RelayGetPower) PayloadType() uint16 {
	return uint16(PayloadTypeRelayGetPower)
}

// Size is the total size of the message.
func (p *RelayGetPower) Size() int {
	return 1
}

// MarshalBinary encodes the RelayGetPower packet into LIFX binary format.
func (p *RelayGetPower) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.RelayIndex); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the RelayGetPower struct.
func (p *RelayGetPower) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.RelayIndex); err != nil {
		return err
	}
	return nil
}

// RelaySetPower defines the fields for a packet of 3 bytes.
type RelaySetPower struct {
	RelayIndex uint8
	Level      uint16
}

// PayloadType returns the LIFX message type for RelaySetPower.
func (p *RelaySetPower) PayloadType() uint16 {
	return uint16(PayloadTypeRelaySetPower)
}

// Size is the total size of the message.
func (p *RelaySetPower) Size() int {
	return 3
}

// MarshalBinary encodes the RelaySetPower packet into LIFX binary format.
func (p *RelaySetPower) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.RelayIndex); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Level); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the RelaySetPower struct.
func (p *RelaySetPower) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.RelayIndex); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Level); err != nil {
		return err
	}
	return nil
}

// RelayStatePower defines the fields for a packet of 3 bytes.
type RelayStatePower struct {
	RelayIndex uint8
	Level      uint16
}

// PayloadType returns the LIFX message type for RelayStatePower.
func (p *RelayStatePower) PayloadType() uint16 {
	return uint16(PayloadTypeRelayStatePower)
}

// Size is the total size of the message.
func (p *RelayStatePower) Size() int {
	return 3
}

// MarshalBinary encodes the RelayStatePower packet into LIFX binary format.
func (p *RelayStatePower) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.RelayIndex); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Level); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the RelayStatePower struct.
func (p *RelayStatePower) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.RelayIndex); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Level); err != nil {
		return err
	}
	return nil
}
