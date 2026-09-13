// Code generated. DO NOT EDIT.
// Source: https://github.com/LIFX/public-protocol@9289015784bd338a546ce4fe67daa1df0b2abfcd
// Generated: 2026-09-13T04:04:50Z
package packets

import (
	"bytes"
	"encoding/binary"
)

// SensorGetAmbientLight defines the fields for a packet of 0 bytes.
type SensorGetAmbientLight struct {
}

// PayloadType returns the LIFX message type for SensorGetAmbientLight.
func (p *SensorGetAmbientLight) PayloadType() uint16 {
	return uint16(PayloadTypeSensorGetAmbientLight)
}

// Size is the total size of the message.
func (p *SensorGetAmbientLight) Size() int {
	return 0
}

// MarshalBinary encodes the SensorGetAmbientLight packet into LIFX binary format.
func (p *SensorGetAmbientLight) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the SensorGetAmbientLight struct.
func (p *SensorGetAmbientLight) UnmarshalBinary(data []byte) error {
	return nil
}

// SensorStateAmbientLight defines the fields for a packet of 4 bytes.
type SensorStateAmbientLight struct {
	Lux float32
}

// PayloadType returns the LIFX message type for SensorStateAmbientLight.
func (p *SensorStateAmbientLight) PayloadType() uint16 {
	return uint16(PayloadTypeSensorStateAmbientLight)
}

// Size is the total size of the message.
func (p *SensorStateAmbientLight) Size() int {
	return 4
}

// MarshalBinary encodes the SensorStateAmbientLight packet into LIFX binary format.
func (p *SensorStateAmbientLight) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Lux); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the SensorStateAmbientLight struct.
func (p *SensorStateAmbientLight) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Lux); err != nil {
		return err
	}
	return nil
}
