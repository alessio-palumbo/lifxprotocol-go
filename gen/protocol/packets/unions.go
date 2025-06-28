// Code generated. DO NOT EDIT.
// Source: https://github.com/LIFX/public-protocol@f061a60a2e9305e63ca75cfa53ad78d276de0e0a
// Generated: 2025-06-28T09:47:36Z
package packets

import "unsafe"

// ButtonTarget is a union that can hold one of several field variants in a fixed-size byte array.
type ButtonTarget [16]byte

// Relays returns the union as a pointer to ButtonTargetRelays.
func (u *ButtonTarget) Relays() *ButtonTargetRelays {
	return (*ButtonTargetRelays)(unsafe.Pointer(u))
}

// SetRelays sets the union using the provided *ButtonTargetRelays.
func (u *ButtonTarget) SetRelays(v *ButtonTargetRelays) {
	if v == nil {
		var zero ButtonTargetRelays
		v = &zero
	}
	copy(u[0:16], (*[16]byte)(unsafe.Pointer(v))[:])
}

// Device returns the union as a pointer to ButtonTargetDevice.
func (u *ButtonTarget) Device() *ButtonTargetDevice {
	return (*ButtonTargetDevice)(unsafe.Pointer(u))
}

// SetDevice sets the union using the provided *ButtonTargetDevice.
func (u *ButtonTarget) SetDevice(v *ButtonTargetDevice) {
	if v == nil {
		var zero ButtonTargetDevice
		v = &zero
	}
	copy(u[0:16], (*[16]byte)(unsafe.Pointer(v))[:])
}

// Location returns the raw bytes of the union for the Location variant.
func (u *ButtonTarget) Location() []byte {
	return u[0:16]
}

// SetLocation sets the union using the provided raw bytes.
func (u *ButtonTarget) SetLocation(v []byte) {
	if len(v) != 16 {
		panic("invalid length for Location: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// Group returns the raw bytes of the union for the Group variant.
func (u *ButtonTarget) Group() []byte {
	return u[0:16]
}

// SetGroup sets the union using the provided raw bytes.
func (u *ButtonTarget) SetGroup(v []byte) {
	if len(v) != 16 {
		panic("invalid length for Group: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// Scene returns the raw bytes of the union for the Scene variant.
func (u *ButtonTarget) Scene() []byte {
	return u[0:16]
}

// SetScene sets the union using the provided raw bytes.
func (u *ButtonTarget) SetScene(v []byte) {
	if len(v) != 16 {
		panic("invalid length for Scene: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// DeviceRelays returns the union as a pointer to ButtonTargetDeviceRelays.
func (u *ButtonTarget) DeviceRelays() *ButtonTargetDeviceRelays {
	return (*ButtonTargetDeviceRelays)(unsafe.Pointer(u))
}

// SetDeviceRelays sets the union using the provided *ButtonTargetDeviceRelays.
func (u *ButtonTarget) SetDeviceRelays(v *ButtonTargetDeviceRelays) {
	if v == nil {
		var zero ButtonTargetDeviceRelays
		v = &zero
	}
	copy(u[0:16], (*[16]byte)(unsafe.Pointer(v))[:])
}
