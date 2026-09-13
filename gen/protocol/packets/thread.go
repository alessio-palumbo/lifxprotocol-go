// Code generated. DO NOT EDIT.
// Source: https://github.com/LIFX/public-protocol@9289015784bd338a546ce4fe67daa1df0b2abfcd
// Generated: 2026-09-13T04:04:50Z
package packets

import (
	"bytes"
	"encoding/binary"

	"github.com/alessio-palumbo/lifxprotocol-go/gen/protocol/enums"
)

// ThreadGetInfo defines the fields for a packet of 0 bytes.
type ThreadGetInfo struct {
}

// PayloadType returns the LIFX message type for ThreadGetInfo.
func (p *ThreadGetInfo) PayloadType() uint16 {
	return uint16(PayloadTypeThreadGetInfo)
}

// Size is the total size of the message.
func (p *ThreadGetInfo) Size() int {
	return 0
}

// MarshalBinary encodes the ThreadGetInfo packet into LIFX binary format.
func (p *ThreadGetInfo) MarshalBinary() ([]byte, error) {
	return nil, nil
}

// UnmarshalBinary decodes the LIFX binary format into the ThreadGetInfo struct.
func (p *ThreadGetInfo) UnmarshalBinary(data []byte) error {
	return nil
}

// ThreadStateInfo defines the fields for a packet of 32 bytes.
type ThreadStateInfo struct {
	Rloc16      uint16
	Reserved1   uint16
	NetworkName [16]byte
	Role        enums.ThreadRoutingRole
	Reserved2   uint8
	Reserved3   uint8
	Reserved4   uint8
	LinkHealth  ThreadLinkHealth
}

// PayloadType returns the LIFX message type for ThreadStateInfo.
func (p *ThreadStateInfo) PayloadType() uint16 {
	return uint16(PayloadTypeThreadStateInfo)
}

// Size is the total size of the message.
func (p *ThreadStateInfo) Size() int {
	return 32
}

// MarshalBinary encodes the ThreadStateInfo packet into LIFX binary format.
func (p *ThreadStateInfo) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Rloc16); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved1); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.NetworkName); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Role); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved2); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved3); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved4); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.LinkHealth); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the ThreadStateInfo struct.
func (p *ThreadStateInfo) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Rloc16); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved1); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.NetworkName); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Role); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved2); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved3); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved4); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.LinkHealth); err != nil {
		return err
	}
	return nil
}
