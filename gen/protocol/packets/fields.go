// Code generated. DO NOT EDIT.
// Source: https://github.com/LIFX/public-protocol@f061a60a2e9305e63ca75cfa53ad78d276de0e0a
// Generated: 2025-06-28T09:47:36Z
package packets

import "github.com/alessio-palumbo/lifxprotocol-go/gen/protocol/enums"

// ButtonTargetRelays defines the fields for a group of 16 bytes.
type ButtonTargetRelays struct {
	RelaysCount uint8
	Relays      [15]uint8
}

// ButtonTargetDevice defines the fields for a group of 16 bytes.
type ButtonTargetDevice struct {
	Serial   [6]byte
	Reserved [10]byte
}

// ButtonTargetDeviceRelays defines the fields for a group of 16 bytes.
type ButtonTargetDeviceRelays struct {
	Serial      [6]byte
	RelaysCount uint8
	Relays      [9]uint8
}

// ButtonAction defines the fields for a group of 20 bytes.
type ButtonAction struct {
	Gesture    enums.ButtonGesture
	TargetType enums.ButtonTargetType
	Target     ButtonTarget
}

// Button defines the fields for a group of 101 bytes.
type Button struct {
	ActionsCount uint8
	Actions      [5]ButtonAction
}

// ButtonBacklightHsbk defines the fields for a group of 8 bytes.
type ButtonBacklightHsbk struct {
	Hue        uint16
	Saturation uint16
	Brightness uint16
	Kelvin     uint16
}

// LightHsbk defines the fields for a group of 8 bytes.
type LightHsbk struct {
	Hue        uint16
	Saturation uint16
	Brightness uint16
	Kelvin     uint16
}

// MultiZoneEffectParameter defines the fields for a group of 32 bytes.
type MultiZoneEffectParameter struct {
	Parameter0 uint32
	Parameter1 uint32
	Parameter2 uint32
	Parameter3 uint32
	Parameter4 uint32
	Parameter5 uint32
	Parameter6 uint32
	Parameter7 uint32
}

// MultiZoneEffectSettings defines the fields for a group of 59 bytes.
type MultiZoneEffectSettings struct {
	Instanceid uint32
	Type       enums.MultiZoneEffectType
	Reserved1  uint16
	Speed      uint32
	Duration   uint64
	Reserved2  uint32
	Reserved3  uint32
	Parameter  MultiZoneEffectParameter
}

// TileAccelMeas defines the fields for a group of 6 bytes.
type TileAccelMeas struct {
	X int16
	Y int16
	Z int16
}

// TileStateDevice defines the fields for a group of 55 bytes.
type TileStateDevice struct {
	AccelMeas     TileAccelMeas
	Reserved1     uint8
	Reserved2     uint8
	UserX         float32
	UserY         float32
	Width         uint8
	Height        uint8
	Reserved3     uint8
	DeviceVersion DeviceStateVersion
	Firmware      DeviceStateHostFirmware
	Reserved4     uint32
}

// TileBufferRect defines the fields for a group of 4 bytes.
type TileBufferRect struct {
	Reserved1 uint8
	X         uint8
	Y         uint8
	Width     uint8
}

// TileEffectParameter defines the fields for a group of 32 bytes.
type TileEffectParameter struct {
	Parameter0 uint32
	Parameter1 uint32
	Parameter2 uint32
	Parameter3 uint32
	Parameter4 uint32
	Parameter5 uint32
	Parameter6 uint32
	Parameter7 uint32
}

// TileEffectSettings defines the fields for a group of 186 bytes.
type TileEffectSettings struct {
	Instanceid   uint32
	Type         enums.TileEffectType
	Speed        uint32
	Duration     uint64
	Reserved1    uint32
	Reserved2    uint32
	Parameter    TileEffectParameter
	PaletteCount uint8
	Palette      [16]LightHsbk
}
