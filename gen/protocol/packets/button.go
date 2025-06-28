// Code generated. DO NOT EDIT.
// Source: https://github.com/LIFX/public-protocol@f061a60a2e9305e63ca75cfa53ad78d276de0e0a
// Generated: 2025-06-28T09:47:36Z
package packets

import (
	"bytes"
	"encoding/binary"
)

// ButtonGet defines the fields for a packet of 0 bytes.
type ButtonGet struct {
}

// PayloadType returns the LIFX message type for ButtonGet.
func (p *ButtonGet) PayloadType() uint16 {
	return uint16(PayloadTypeButtonGet)
}

// Size is the total size of the message.
func (p *ButtonGet) Size() int {
	return 0
}

// MarshalBinary encodes the ButtonGet packet into LIFX binary format.
func (p *ButtonGet) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the ButtonGet struct.
func (p *ButtonGet) UnmarshalBinary(data []byte) error {
	return nil
}

// ButtonSet defines the fields for a packet of 810 bytes.
type ButtonSet struct {
	Index        uint8
	ButtonsCount uint8
	Buttons      [8]Button
}

// PayloadType returns the LIFX message type for ButtonSet.
func (p *ButtonSet) PayloadType() uint16 {
	return uint16(PayloadTypeButtonSet)
}

// Size is the total size of the message.
func (p *ButtonSet) Size() int {
	return 810
}

// MarshalBinary encodes the ButtonSet packet into LIFX binary format.
func (p *ButtonSet) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Index); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.ButtonsCount); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Buttons); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the ButtonSet struct.
func (p *ButtonSet) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Index); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.ButtonsCount); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Buttons); err != nil {
		return err
	}
	return nil
}

// ButtonState defines the fields for a packet of 811 bytes.
type ButtonState struct {
	Count        uint8
	Index        uint8
	ButtonsCount uint8
	Buttons      [8]Button
}

// PayloadType returns the LIFX message type for ButtonState.
func (p *ButtonState) PayloadType() uint16 {
	return uint16(PayloadTypeButtonState)
}

// Size is the total size of the message.
func (p *ButtonState) Size() int {
	return 811
}

// MarshalBinary encodes the ButtonState packet into LIFX binary format.
func (p *ButtonState) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Count); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Index); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.ButtonsCount); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Buttons); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the ButtonState struct.
func (p *ButtonState) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Count); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Index); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.ButtonsCount); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Buttons); err != nil {
		return err
	}
	return nil
}

// ButtonGetConfig defines the fields for a packet of 0 bytes.
type ButtonGetConfig struct {
}

// PayloadType returns the LIFX message type for ButtonGetConfig.
func (p *ButtonGetConfig) PayloadType() uint16 {
	return uint16(PayloadTypeButtonGetConfig)
}

// Size is the total size of the message.
func (p *ButtonGetConfig) Size() int {
	return 0
}

// MarshalBinary encodes the ButtonGetConfig packet into LIFX binary format.
func (p *ButtonGetConfig) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the ButtonGetConfig struct.
func (p *ButtonGetConfig) UnmarshalBinary(data []byte) error {
	return nil
}

// ButtonStateConfig defines the fields for a packet of 18 bytes.
type ButtonStateConfig struct {
	HapticDurationMs  uint16
	BacklightOnColor  ButtonBacklightHsbk
	BacklightOffColor ButtonBacklightHsbk
}

// PayloadType returns the LIFX message type for ButtonStateConfig.
func (p *ButtonStateConfig) PayloadType() uint16 {
	return uint16(PayloadTypeButtonStateConfig)
}

// Size is the total size of the message.
func (p *ButtonStateConfig) Size() int {
	return 18
}

// MarshalBinary encodes the ButtonStateConfig packet into LIFX binary format.
func (p *ButtonStateConfig) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.HapticDurationMs); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.BacklightOnColor); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.BacklightOffColor); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the ButtonStateConfig struct.
func (p *ButtonStateConfig) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.HapticDurationMs); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.BacklightOnColor); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.BacklightOffColor); err != nil {
		return err
	}
	return nil
}

// ButtonSetConfig defines the fields for a packet of 18 bytes.
type ButtonSetConfig struct {
	HapticDurationMs  uint16
	BacklightOnColor  ButtonBacklightHsbk
	BacklightOffColor ButtonBacklightHsbk
}

// PayloadType returns the LIFX message type for ButtonSetConfig.
func (p *ButtonSetConfig) PayloadType() uint16 {
	return uint16(PayloadTypeButtonSetConfig)
}

// Size is the total size of the message.
func (p *ButtonSetConfig) Size() int {
	return 18
}

// MarshalBinary encodes the ButtonSetConfig packet into LIFX binary format.
func (p *ButtonSetConfig) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.HapticDurationMs); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.BacklightOnColor); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.BacklightOffColor); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the ButtonSetConfig struct.
func (p *ButtonSetConfig) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.HapticDurationMs); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.BacklightOnColor); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.BacklightOffColor); err != nil {
		return err
	}
	return nil
}
