// Code generated. DO NOT EDIT.
package packets

import (
	"bytes"
	"encoding/binary"
)

// TestPacket defines the fields for a packet of 16 bytes.
type TestPacket struct {
	Serial    [6]byte
	Reserved1 [10]byte
}

// PayloadType returns the LIFX message type for TestPacket.
func (p *TestPacket) PayloadType() uint16 {
	return uint16(PayloadTypeTestPacket)
}

// Size is the total size of the message.
func (p *TestPacket) Size() int {
	return 16
}

// MarshalBinary encodes the TestPacket packet into LIFX binary format.
func (p *TestPacket) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, p.Serial); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.LittleEndian, p.Reserved1); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary decodes the LIFX binary format into the TestPacket struct.
func (p *TestPacket) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &p.Serial); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.LittleEndian, &p.Reserved1); err != nil {
		return err
	}
	return nil
}
