// Code generated. DO NOT EDIT.
// Source: https://github.com/LIFX/public-protocol@f061a60a2e9305e63ca75cfa53ad78d276de0e0a
// Generated: 2025-06-28T09:47:36Z
package packets

import (
	"bytes"
	"encoding/binary"

	"github.com/alessio-palumbo/lifxprotocol-go/gen/protocol/enums"
)

// LightGet defines the fields for a packet of 0 bytes.
type LightGet struct {
}

// PayloadType returns the LIFX message type for LightGet.
func (p *LightGet) PayloadType() uint16 {
	return uint16(PayloadTypeLightGet)
}

// Size is the total size of the message.
func (p *LightGet) Size() int {
	return 0
}

// MarshalBinary encodes the LightGet packet into LIFX binary format.
func (p *LightGet) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightGet struct.
func (p *LightGet) UnmarshalBinary(data []byte) error {
	return nil
}

// LightSetColor defines the fields for a packet of 13 bytes.
type LightSetColor struct {
	Reserved1 uint8
	Color     LightHsbk
	Duration  uint32
}

// PayloadType returns the LIFX message type for LightSetColor.
func (p *LightSetColor) PayloadType() uint16 {
	return uint16(PayloadTypeLightSetColor)
}

// Size is the total size of the message.
func (p *LightSetColor) Size() int {
	return 13
}

// MarshalBinary encodes the LightSetColor packet into LIFX binary format.
func (p *LightSetColor) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved1); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Color); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Duration); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightSetColor struct.
func (p *LightSetColor) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved1); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Color); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Duration); err != nil {
		return err
	}
	return nil
}

// LightSetWaveformOptional defines the fields for a packet of 25 bytes.
type LightSetWaveformOptional struct {
	Reserved1     uint8
	Transient     bool
	Color         LightHsbk
	Period        uint32
	Cycles        float32
	SkewRatio     int16
	Waveform      enums.LightWaveform
	SetHue        bool
	SetSaturation bool
	SetBrightness bool
	SetKelvin     bool
}

// PayloadType returns the LIFX message type for LightSetWaveformOptional.
func (p *LightSetWaveformOptional) PayloadType() uint16 {
	return uint16(PayloadTypeLightSetWaveformOptional)
}

// Size is the total size of the message.
func (p *LightSetWaveformOptional) Size() int {
	return 25
}

// MarshalBinary encodes the LightSetWaveformOptional packet into LIFX binary format.
func (p *LightSetWaveformOptional) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved1); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, boolToUint8(p.Transient)); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Color); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Period); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Cycles); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.SkewRatio); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Waveform); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, boolToUint8(p.SetHue)); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, boolToUint8(p.SetSaturation)); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, boolToUint8(p.SetBrightness)); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, boolToUint8(p.SetKelvin)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightSetWaveformOptional struct.
func (p *LightSetWaveformOptional) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved1); err != nil {
		return err
	}
	{
		var b uint8
		if err := binary.Read(buf, binary.LittleEndian, &b); err != nil {
			return err
		}
		p.Transient = b != 0
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Color); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Period); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Cycles); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.SkewRatio); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Waveform); err != nil {
		return err
	}
	{
		var b uint8
		if err := binary.Read(buf, binary.LittleEndian, &b); err != nil {
			return err
		}
		p.SetHue = b != 0
	}
	{
		var b uint8
		if err := binary.Read(buf, binary.LittleEndian, &b); err != nil {
			return err
		}
		p.SetSaturation = b != 0
	}
	{
		var b uint8
		if err := binary.Read(buf, binary.LittleEndian, &b); err != nil {
			return err
		}
		p.SetBrightness = b != 0
	}
	{
		var b uint8
		if err := binary.Read(buf, binary.LittleEndian, &b); err != nil {
			return err
		}
		p.SetKelvin = b != 0
	}
	return nil
}

// LightSetWaveform defines the fields for a packet of 21 bytes.
type LightSetWaveform struct {
	Reserved1 uint8
	Transient bool
	Color     LightHsbk
	Period    uint32
	Cycles    float32
	SkewRatio int16
	Waveform  enums.LightWaveform
}

// PayloadType returns the LIFX message type for LightSetWaveform.
func (p *LightSetWaveform) PayloadType() uint16 {
	return uint16(PayloadTypeLightSetWaveform)
}

// Size is the total size of the message.
func (p *LightSetWaveform) Size() int {
	return 21
}

// MarshalBinary encodes the LightSetWaveform packet into LIFX binary format.
func (p *LightSetWaveform) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved1); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, boolToUint8(p.Transient)); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Color); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Period); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Cycles); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.SkewRatio); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Waveform); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightSetWaveform struct.
func (p *LightSetWaveform) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved1); err != nil {
		return err
	}
	{
		var b uint8
		if err := binary.Read(buf, binary.LittleEndian, &b); err != nil {
			return err
		}
		p.Transient = b != 0
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Color); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Period); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Cycles); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.SkewRatio); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Waveform); err != nil {
		return err
	}
	return nil
}

// LightGetPower defines the fields for a packet of 0 bytes.
type LightGetPower struct {
}

// PayloadType returns the LIFX message type for LightGetPower.
func (p *LightGetPower) PayloadType() uint16 {
	return uint16(PayloadTypeLightGetPower)
}

// Size is the total size of the message.
func (p *LightGetPower) Size() int {
	return 0
}

// MarshalBinary encodes the LightGetPower packet into LIFX binary format.
func (p *LightGetPower) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightGetPower struct.
func (p *LightGetPower) UnmarshalBinary(data []byte) error {
	return nil
}

// LightSetPower defines the fields for a packet of 6 bytes.
type LightSetPower struct {
	Level    uint16
	Duration uint32
}

// PayloadType returns the LIFX message type for LightSetPower.
func (p *LightSetPower) PayloadType() uint16 {
	return uint16(PayloadTypeLightSetPower)
}

// Size is the total size of the message.
func (p *LightSetPower) Size() int {
	return 6
}

// MarshalBinary encodes the LightSetPower packet into LIFX binary format.
func (p *LightSetPower) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Level); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Duration); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightSetPower struct.
func (p *LightSetPower) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Level); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Duration); err != nil {
		return err
	}
	return nil
}

// LightStatePower defines the fields for a packet of 2 bytes.
type LightStatePower struct {
	Level uint16
}

// PayloadType returns the LIFX message type for LightStatePower.
func (p *LightStatePower) PayloadType() uint16 {
	return uint16(PayloadTypeLightStatePower)
}

// Size is the total size of the message.
func (p *LightStatePower) Size() int {
	return 2
}

// MarshalBinary encodes the LightStatePower packet into LIFX binary format.
func (p *LightStatePower) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Level); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightStatePower struct.
func (p *LightStatePower) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Level); err != nil {
		return err
	}
	return nil
}

// LightState defines the fields for a packet of 52 bytes.
type LightState struct {
	Color     LightHsbk
	Reserved1 uint16
	Power     uint16
	Label     [32]byte
	Reserved2 uint64
}

// PayloadType returns the LIFX message type for LightState.
func (p *LightState) PayloadType() uint16 {
	return uint16(PayloadTypeLightState)
}

// Size is the total size of the message.
func (p *LightState) Size() int {
	return 52
}

// MarshalBinary encodes the LightState packet into LIFX binary format.
func (p *LightState) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Color); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved1); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Power); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Label); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved2); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightState struct.
func (p *LightState) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Color); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved1); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Power); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Label); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved2); err != nil {
		return err
	}
	return nil
}

// LightGetInfrared defines the fields for a packet of 0 bytes.
type LightGetInfrared struct {
}

// PayloadType returns the LIFX message type for LightGetInfrared.
func (p *LightGetInfrared) PayloadType() uint16 {
	return uint16(PayloadTypeLightGetInfrared)
}

// Size is the total size of the message.
func (p *LightGetInfrared) Size() int {
	return 0
}

// MarshalBinary encodes the LightGetInfrared packet into LIFX binary format.
func (p *LightGetInfrared) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightGetInfrared struct.
func (p *LightGetInfrared) UnmarshalBinary(data []byte) error {
	return nil
}

// LightStateInfrared defines the fields for a packet of 2 bytes.
type LightStateInfrared struct {
	Brightness uint16
}

// PayloadType returns the LIFX message type for LightStateInfrared.
func (p *LightStateInfrared) PayloadType() uint16 {
	return uint16(PayloadTypeLightStateInfrared)
}

// Size is the total size of the message.
func (p *LightStateInfrared) Size() int {
	return 2
}

// MarshalBinary encodes the LightStateInfrared packet into LIFX binary format.
func (p *LightStateInfrared) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Brightness); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightStateInfrared struct.
func (p *LightStateInfrared) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Brightness); err != nil {
		return err
	}
	return nil
}

// LightSetInfrared defines the fields for a packet of 2 bytes.
type LightSetInfrared struct {
	Brightness uint16
}

// PayloadType returns the LIFX message type for LightSetInfrared.
func (p *LightSetInfrared) PayloadType() uint16 {
	return uint16(PayloadTypeLightSetInfrared)
}

// Size is the total size of the message.
func (p *LightSetInfrared) Size() int {
	return 2
}

// MarshalBinary encodes the LightSetInfrared packet into LIFX binary format.
func (p *LightSetInfrared) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Brightness); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightSetInfrared struct.
func (p *LightSetInfrared) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Brightness); err != nil {
		return err
	}
	return nil
}

// LightGetHevCycle defines the fields for a packet of 0 bytes.
type LightGetHevCycle struct {
}

// PayloadType returns the LIFX message type for LightGetHevCycle.
func (p *LightGetHevCycle) PayloadType() uint16 {
	return uint16(PayloadTypeLightGetHevCycle)
}

// Size is the total size of the message.
func (p *LightGetHevCycle) Size() int {
	return 0
}

// MarshalBinary encodes the LightGetHevCycle packet into LIFX binary format.
func (p *LightGetHevCycle) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightGetHevCycle struct.
func (p *LightGetHevCycle) UnmarshalBinary(data []byte) error {
	return nil
}

// LightSetHevCycle defines the fields for a packet of 5 bytes.
type LightSetHevCycle struct {
	Enable    bool
	DurationS uint32
}

// PayloadType returns the LIFX message type for LightSetHevCycle.
func (p *LightSetHevCycle) PayloadType() uint16 {
	return uint16(PayloadTypeLightSetHevCycle)
}

// Size is the total size of the message.
func (p *LightSetHevCycle) Size() int {
	return 5
}

// MarshalBinary encodes the LightSetHevCycle packet into LIFX binary format.
func (p *LightSetHevCycle) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, boolToUint8(p.Enable)); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.DurationS); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightSetHevCycle struct.
func (p *LightSetHevCycle) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	{
		var b uint8
		if err := binary.Read(buf, binary.LittleEndian, &b); err != nil {
			return err
		}
		p.Enable = b != 0
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.DurationS); err != nil {
		return err
	}
	return nil
}

// LightStateHevCycle defines the fields for a packet of 9 bytes.
type LightStateHevCycle struct {
	DurationS  uint32
	RemainingS uint32
	LastPower  bool
}

// PayloadType returns the LIFX message type for LightStateHevCycle.
func (p *LightStateHevCycle) PayloadType() uint16 {
	return uint16(PayloadTypeLightStateHevCycle)
}

// Size is the total size of the message.
func (p *LightStateHevCycle) Size() int {
	return 9
}

// MarshalBinary encodes the LightStateHevCycle packet into LIFX binary format.
func (p *LightStateHevCycle) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.DurationS); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.RemainingS); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, boolToUint8(p.LastPower)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightStateHevCycle struct.
func (p *LightStateHevCycle) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.DurationS); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.RemainingS); err != nil {
		return err
	}
	{
		var b uint8
		if err := binary.Read(buf, binary.LittleEndian, &b); err != nil {
			return err
		}
		p.LastPower = b != 0
	}
	return nil
}

// LightGetHevCycleConfiguration defines the fields for a packet of 0 bytes.
type LightGetHevCycleConfiguration struct {
}

// PayloadType returns the LIFX message type for LightGetHevCycleConfiguration.
func (p *LightGetHevCycleConfiguration) PayloadType() uint16 {
	return uint16(PayloadTypeLightGetHevCycleConfiguration)
}

// Size is the total size of the message.
func (p *LightGetHevCycleConfiguration) Size() int {
	return 0
}

// MarshalBinary encodes the LightGetHevCycleConfiguration packet into LIFX binary format.
func (p *LightGetHevCycleConfiguration) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightGetHevCycleConfiguration struct.
func (p *LightGetHevCycleConfiguration) UnmarshalBinary(data []byte) error {
	return nil
}

// LightSetHevCycleConfiguration defines the fields for a packet of 5 bytes.
type LightSetHevCycleConfiguration struct {
	Indication bool
	DurationS  uint32
}

// PayloadType returns the LIFX message type for LightSetHevCycleConfiguration.
func (p *LightSetHevCycleConfiguration) PayloadType() uint16 {
	return uint16(PayloadTypeLightSetHevCycleConfiguration)
}

// Size is the total size of the message.
func (p *LightSetHevCycleConfiguration) Size() int {
	return 5
}

// MarshalBinary encodes the LightSetHevCycleConfiguration packet into LIFX binary format.
func (p *LightSetHevCycleConfiguration) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, boolToUint8(p.Indication)); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.DurationS); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightSetHevCycleConfiguration struct.
func (p *LightSetHevCycleConfiguration) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	{
		var b uint8
		if err := binary.Read(buf, binary.LittleEndian, &b); err != nil {
			return err
		}
		p.Indication = b != 0
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.DurationS); err != nil {
		return err
	}
	return nil
}

// LightStateHevCycleConfiguration defines the fields for a packet of 5 bytes.
type LightStateHevCycleConfiguration struct {
	Indication bool
	DurationS  uint32
}

// PayloadType returns the LIFX message type for LightStateHevCycleConfiguration.
func (p *LightStateHevCycleConfiguration) PayloadType() uint16 {
	return uint16(PayloadTypeLightStateHevCycleConfiguration)
}

// Size is the total size of the message.
func (p *LightStateHevCycleConfiguration) Size() int {
	return 5
}

// MarshalBinary encodes the LightStateHevCycleConfiguration packet into LIFX binary format.
func (p *LightStateHevCycleConfiguration) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, boolToUint8(p.Indication)); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.DurationS); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightStateHevCycleConfiguration struct.
func (p *LightStateHevCycleConfiguration) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	{
		var b uint8
		if err := binary.Read(buf, binary.LittleEndian, &b); err != nil {
			return err
		}
		p.Indication = b != 0
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.DurationS); err != nil {
		return err
	}
	return nil
}

// LightGetLastHevCycleResult defines the fields for a packet of 0 bytes.
type LightGetLastHevCycleResult struct {
}

// PayloadType returns the LIFX message type for LightGetLastHevCycleResult.
func (p *LightGetLastHevCycleResult) PayloadType() uint16 {
	return uint16(PayloadTypeLightGetLastHevCycleResult)
}

// Size is the total size of the message.
func (p *LightGetLastHevCycleResult) Size() int {
	return 0
}

// MarshalBinary encodes the LightGetLastHevCycleResult packet into LIFX binary format.
func (p *LightGetLastHevCycleResult) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightGetLastHevCycleResult struct.
func (p *LightGetLastHevCycleResult) UnmarshalBinary(data []byte) error {
	return nil
}

// LightStateLastHevCycleResult defines the fields for a packet of 1 bytes.
type LightStateLastHevCycleResult struct {
	Result enums.LightLastHevCycleResult
}

// PayloadType returns the LIFX message type for LightStateLastHevCycleResult.
func (p *LightStateLastHevCycleResult) PayloadType() uint16 {
	return uint16(PayloadTypeLightStateLastHevCycleResult)
}

// Size is the total size of the message.
func (p *LightStateLastHevCycleResult) Size() int {
	return 1
}

// MarshalBinary encodes the LightStateLastHevCycleResult packet into LIFX binary format.
func (p *LightStateLastHevCycleResult) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Result); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the LightStateLastHevCycleResult struct.
func (p *LightStateLastHevCycleResult) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Result); err != nil {
		return err
	}
	return nil
}
