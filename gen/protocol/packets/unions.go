// Code generated. DO NOT EDIT.
// Source: https://github.com/LIFX/public-protocol@fe8871b049893401c8c6a5fc425971ec55a2fa06
// Generated: 2025-12-17T01:06:02Z
package packets

import "unsafe"

// ButtonTarget is a union that can hold one of several field variants in a fixed-size byte array.
type ButtonTarget [16]byte

// PowerToggleRelays returns the union as a pointer to ButtonTargetRelays.
func (u *ButtonTarget) PowerToggleRelays() *ButtonTargetRelays {
	return (*ButtonTargetRelays)(unsafe.Pointer(u))
}

// SetPowerToggleRelays sets the union using the provided *ButtonTargetRelays.
func (u *ButtonTarget) SetPowerToggleRelays(v *ButtonTargetRelays) {
	if v == nil {
		var zero ButtonTargetRelays
		v = &zero
	}
	copy(u[0:16], (*[16]byte)(unsafe.Pointer(v))[:])
}

// PowerToggleDevice returns the union as a pointer to ButtonTargetDevice.
func (u *ButtonTarget) PowerToggleDevice() *ButtonTargetDevice {
	return (*ButtonTargetDevice)(unsafe.Pointer(u))
}

// SetPowerToggleDevice sets the union using the provided *ButtonTargetDevice.
func (u *ButtonTarget) SetPowerToggleDevice(v *ButtonTargetDevice) {
	if v == nil {
		var zero ButtonTargetDevice
		v = &zero
	}
	copy(u[0:16], (*[16]byte)(unsafe.Pointer(v))[:])
}

// PowerToggleLocation returns the raw bytes of the union for the PowerToggleLocation variant.
func (u *ButtonTarget) PowerToggleLocation() []byte {
	return u[0:16]
}

// SetPowerToggleLocation sets the union using the provided raw bytes.
func (u *ButtonTarget) SetPowerToggleLocation(v []byte) {
	if len(v) != 16 {
		panic("invalid length for PowerToggleLocation: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// PowerToggleGroup returns the raw bytes of the union for the PowerToggleGroup variant.
func (u *ButtonTarget) PowerToggleGroup() []byte {
	return u[0:16]
}

// SetPowerToggleGroup sets the union using the provided raw bytes.
func (u *ButtonTarget) SetPowerToggleGroup(v []byte) {
	if len(v) != 16 {
		panic("invalid length for PowerToggleGroup: expected 16 bytes")
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

// PowerToggleDeviceRelays returns the union as a pointer to ButtonTargetDeviceRelays.
func (u *ButtonTarget) PowerToggleDeviceRelays() *ButtonTargetDeviceRelays {
	return (*ButtonTargetDeviceRelays)(unsafe.Pointer(u))
}

// SetPowerToggleDeviceRelays sets the union using the provided *ButtonTargetDeviceRelays.
func (u *ButtonTarget) SetPowerToggleDeviceRelays(v *ButtonTargetDeviceRelays) {
	if v == nil {
		var zero ButtonTargetDeviceRelays
		v = &zero
	}
	copy(u[0:16], (*[16]byte)(unsafe.Pointer(v))[:])
}

// BrightnessDownDevice returns the union as a pointer to ButtonTargetDevice.
func (u *ButtonTarget) BrightnessDownDevice() *ButtonTargetDevice {
	return (*ButtonTargetDevice)(unsafe.Pointer(u))
}

// SetBrightnessDownDevice sets the union using the provided *ButtonTargetDevice.
func (u *ButtonTarget) SetBrightnessDownDevice(v *ButtonTargetDevice) {
	if v == nil {
		var zero ButtonTargetDevice
		v = &zero
	}
	copy(u[0:16], (*[16]byte)(unsafe.Pointer(v))[:])
}

// BrightnessDownGroup returns the raw bytes of the union for the BrightnessDownGroup variant.
func (u *ButtonTarget) BrightnessDownGroup() []byte {
	return u[0:16]
}

// SetBrightnessDownGroup sets the union using the provided raw bytes.
func (u *ButtonTarget) SetBrightnessDownGroup(v []byte) {
	if len(v) != 16 {
		panic("invalid length for BrightnessDownGroup: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// BrightnessDownLocation returns the raw bytes of the union for the BrightnessDownLocation variant.
func (u *ButtonTarget) BrightnessDownLocation() []byte {
	return u[0:16]
}

// SetBrightnessDownLocation sets the union using the provided raw bytes.
func (u *ButtonTarget) SetBrightnessDownLocation(v []byte) {
	if len(v) != 16 {
		panic("invalid length for BrightnessDownLocation: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// BrightnessUpDevice returns the union as a pointer to ButtonTargetDevice.
func (u *ButtonTarget) BrightnessUpDevice() *ButtonTargetDevice {
	return (*ButtonTargetDevice)(unsafe.Pointer(u))
}

// SetBrightnessUpDevice sets the union using the provided *ButtonTargetDevice.
func (u *ButtonTarget) SetBrightnessUpDevice(v *ButtonTargetDevice) {
	if v == nil {
		var zero ButtonTargetDevice
		v = &zero
	}
	copy(u[0:16], (*[16]byte)(unsafe.Pointer(v))[:])
}

// BrightnessUpGroup returns the raw bytes of the union for the BrightnessUpGroup variant.
func (u *ButtonTarget) BrightnessUpGroup() []byte {
	return u[0:16]
}

// SetBrightnessUpGroup sets the union using the provided raw bytes.
func (u *ButtonTarget) SetBrightnessUpGroup(v []byte) {
	if len(v) != 16 {
		panic("invalid length for BrightnessUpGroup: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// BrightnessUpLocation returns the raw bytes of the union for the BrightnessUpLocation variant.
func (u *ButtonTarget) BrightnessUpLocation() []byte {
	return u[0:16]
}

// SetBrightnessUpLocation sets the union using the provided raw bytes.
func (u *ButtonTarget) SetBrightnessUpLocation(v []byte) {
	if len(v) != 16 {
		panic("invalid length for BrightnessUpLocation: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// DemoEffectCycle returns the raw bytes of the union for the DemoEffectCycle variant.
func (u *ButtonTarget) DemoEffectCycle() []byte {
	return u[0:16]
}

// SetDemoEffectCycle sets the union using the provided raw bytes.
func (u *ButtonTarget) SetDemoEffectCycle(v []byte) {
	if len(v) != 16 {
		panic("invalid length for DemoEffectCycle: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// DemoEffectCycleStop returns the raw bytes of the union for the DemoEffectCycleStop variant.
func (u *ButtonTarget) DemoEffectCycleStop() []byte {
	return u[0:16]
}

// SetDemoEffectCycleStop sets the union using the provided raw bytes.
func (u *ButtonTarget) SetDemoEffectCycleStop(v []byte) {
	if len(v) != 16 {
		panic("invalid length for DemoEffectCycleStop: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// PowerOnDevice returns the union as a pointer to ButtonTargetDevice.
func (u *ButtonTarget) PowerOnDevice() *ButtonTargetDevice {
	return (*ButtonTargetDevice)(unsafe.Pointer(u))
}

// SetPowerOnDevice sets the union using the provided *ButtonTargetDevice.
func (u *ButtonTarget) SetPowerOnDevice(v *ButtonTargetDevice) {
	if v == nil {
		var zero ButtonTargetDevice
		v = &zero
	}
	copy(u[0:16], (*[16]byte)(unsafe.Pointer(v))[:])
}

// PowerOnLocation returns the raw bytes of the union for the PowerOnLocation variant.
func (u *ButtonTarget) PowerOnLocation() []byte {
	return u[0:16]
}

// SetPowerOnLocation sets the union using the provided raw bytes.
func (u *ButtonTarget) SetPowerOnLocation(v []byte) {
	if len(v) != 16 {
		panic("invalid length for PowerOnLocation: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// PowerOnGroup returns the raw bytes of the union for the PowerOnGroup variant.
func (u *ButtonTarget) PowerOnGroup() []byte {
	return u[0:16]
}

// SetPowerOnGroup sets the union using the provided raw bytes.
func (u *ButtonTarget) SetPowerOnGroup(v []byte) {
	if len(v) != 16 {
		panic("invalid length for PowerOnGroup: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// PowerOnRelays returns the union as a pointer to ButtonTargetDeviceRelays.
func (u *ButtonTarget) PowerOnRelays() *ButtonTargetDeviceRelays {
	return (*ButtonTargetDeviceRelays)(unsafe.Pointer(u))
}

// SetPowerOnRelays sets the union using the provided *ButtonTargetDeviceRelays.
func (u *ButtonTarget) SetPowerOnRelays(v *ButtonTargetDeviceRelays) {
	if v == nil {
		var zero ButtonTargetDeviceRelays
		v = &zero
	}
	copy(u[0:16], (*[16]byte)(unsafe.Pointer(v))[:])
}

// PowerOffDevice returns the union as a pointer to ButtonTargetDevice.
func (u *ButtonTarget) PowerOffDevice() *ButtonTargetDevice {
	return (*ButtonTargetDevice)(unsafe.Pointer(u))
}

// SetPowerOffDevice sets the union using the provided *ButtonTargetDevice.
func (u *ButtonTarget) SetPowerOffDevice(v *ButtonTargetDevice) {
	if v == nil {
		var zero ButtonTargetDevice
		v = &zero
	}
	copy(u[0:16], (*[16]byte)(unsafe.Pointer(v))[:])
}

// PowerOffLocation returns the raw bytes of the union for the PowerOffLocation variant.
func (u *ButtonTarget) PowerOffLocation() []byte {
	return u[0:16]
}

// SetPowerOffLocation sets the union using the provided raw bytes.
func (u *ButtonTarget) SetPowerOffLocation(v []byte) {
	if len(v) != 16 {
		panic("invalid length for PowerOffLocation: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// PowerOffGroup returns the raw bytes of the union for the PowerOffGroup variant.
func (u *ButtonTarget) PowerOffGroup() []byte {
	return u[0:16]
}

// SetPowerOffGroup sets the union using the provided raw bytes.
func (u *ButtonTarget) SetPowerOffGroup(v []byte) {
	if len(v) != 16 {
		panic("invalid length for PowerOffGroup: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// PowerOffRelays returns the union as a pointer to ButtonTargetDeviceRelays.
func (u *ButtonTarget) PowerOffRelays() *ButtonTargetDeviceRelays {
	return (*ButtonTargetDeviceRelays)(unsafe.Pointer(u))
}

// SetPowerOffRelays sets the union using the provided *ButtonTargetDeviceRelays.
func (u *ButtonTarget) SetPowerOffRelays(v *ButtonTargetDeviceRelays) {
	if v == nil {
		var zero ButtonTargetDeviceRelays
		v = &zero
	}
	copy(u[0:16], (*[16]byte)(unsafe.Pointer(v))[:])
}

// PowerToggleLocalDevice returns the raw bytes of the union for the PowerToggleLocalDevice variant.
func (u *ButtonTarget) PowerToggleLocalDevice() []byte {
	return u[0:16]
}

// SetPowerToggleLocalDevice sets the union using the provided raw bytes.
func (u *ButtonTarget) SetPowerToggleLocalDevice(v []byte) {
	if len(v) != 16 {
		panic("invalid length for PowerToggleLocalDevice: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// BrightnessDownLocalDevice returns the raw bytes of the union for the BrightnessDownLocalDevice variant.
func (u *ButtonTarget) BrightnessDownLocalDevice() []byte {
	return u[0:16]
}

// SetBrightnessDownLocalDevice sets the union using the provided raw bytes.
func (u *ButtonTarget) SetBrightnessDownLocalDevice(v []byte) {
	if len(v) != 16 {
		panic("invalid length for BrightnessDownLocalDevice: expected 16 bytes")
	}
	copy(u[0:16], v)
}

// BrightnessUpLocalDevice returns the raw bytes of the union for the BrightnessUpLocalDevice variant.
func (u *ButtonTarget) BrightnessUpLocalDevice() []byte {
	return u[0:16]
}

// SetBrightnessUpLocalDevice sets the union using the provided raw bytes.
func (u *ButtonTarget) SetBrightnessUpLocalDevice(v []byte) {
	if len(v) != 16 {
		panic("invalid length for BrightnessUpLocalDevice: expected 16 bytes")
	}
	copy(u[0:16], v)
}
