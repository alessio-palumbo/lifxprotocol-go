// Code generated. DO NOT EDIT.
package packets

import "unsafe"

// TestUnion is a union that can hold one of several field variants in a fixed-size byte array.
type TestUnion [16]byte

// Serial returns the union as a pointer to [6]byte.
func (u *TestUnion) Serial() *[6]byte {
	return (*[6]byte)(unsafe.Pointer(u))
}

// SetSerial sets the union using the provided *[6]byte.
func (u *TestUnion) SetSerial(v *[6]byte) {
	if v == nil {
		var zero [6]byte
		v = &zero
	}
	copy(u[0:16], (*[16]byte)(unsafe.Pointer(v))[:])
}
