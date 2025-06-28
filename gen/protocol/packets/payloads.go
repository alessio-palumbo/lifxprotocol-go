// Code generated. DO NOT EDIT.
// Source: https://github.com/LIFX/public-protocol@f061a60a2e9305e63ca75cfa53ad78d276de0e0a
// Generated: 2025-06-28T09:47:36Z
package packets

// Payload represents a LIFX packet payload that can be serialized and deserialized.
type Payload interface {
	// MarshalBinary encodes the payload into its binary representation.
	MarshalBinary() ([]byte, error)

	// UnmarshalBinary decodes the binary data into the payload struct.
	UnmarshalBinary([]byte) error

	// PayloadType returns the packet type associated with this payload.
	PayloadType() uint16

	// Size returns the size of the message.
	Size() int
}

type payloadType uint16

const (
	PayloadTypeDeviceGetService                payloadType = 2
	PayloadTypeDeviceStateService              payloadType = 3
	PayloadTypeDeviceGetHostFirmware           payloadType = 14
	PayloadTypeDeviceStateHostFirmware         payloadType = 15
	PayloadTypeDeviceGetWifiInfo               payloadType = 16
	PayloadTypeDeviceStateWifiInfo             payloadType = 17
	PayloadTypeDeviceGetWifiFirmware           payloadType = 18
	PayloadTypeDeviceStateWifiFirmware         payloadType = 19
	PayloadTypeDeviceGetPower                  payloadType = 20
	PayloadTypeDeviceSetPower                  payloadType = 21
	PayloadTypeDeviceStatePower                payloadType = 22
	PayloadTypeDeviceGetLabel                  payloadType = 23
	PayloadTypeDeviceSetLabel                  payloadType = 24
	PayloadTypeDeviceStateLabel                payloadType = 25
	PayloadTypeDeviceGetVersion                payloadType = 32
	PayloadTypeDeviceStateVersion              payloadType = 33
	PayloadTypeDeviceGetInfo                   payloadType = 34
	PayloadTypeDeviceStateInfo                 payloadType = 35
	PayloadTypeDeviceSetReboot                 payloadType = 38
	PayloadTypeDeviceAcknowledgement           payloadType = 45
	PayloadTypeDeviceGetLocation               payloadType = 48
	PayloadTypeDeviceSetLocation               payloadType = 49
	PayloadTypeDeviceStateLocation             payloadType = 50
	PayloadTypeDeviceGetGroup                  payloadType = 51
	PayloadTypeDeviceSetGroup                  payloadType = 52
	PayloadTypeDeviceStateGroup                payloadType = 53
	PayloadTypeDeviceEchoRequest               payloadType = 58
	PayloadTypeDeviceEchoResponse              payloadType = 59
	PayloadTypeLightGet                        payloadType = 101
	PayloadTypeLightSetColor                   payloadType = 102
	PayloadTypeLightSetWaveform                payloadType = 103
	PayloadTypeLightState                      payloadType = 107
	PayloadTypeLightGetPower                   payloadType = 116
	PayloadTypeLightSetPower                   payloadType = 117
	PayloadTypeLightStatePower                 payloadType = 118
	PayloadTypeLightSetWaveformOptional        payloadType = 119
	PayloadTypeLightGetInfrared                payloadType = 120
	PayloadTypeLightStateInfrared              payloadType = 121
	PayloadTypeLightSetInfrared                payloadType = 122
	PayloadTypeLightGetHevCycle                payloadType = 142
	PayloadTypeLightSetHevCycle                payloadType = 143
	PayloadTypeLightStateHevCycle              payloadType = 144
	PayloadTypeLightGetHevCycleConfiguration   payloadType = 145
	PayloadTypeLightSetHevCycleConfiguration   payloadType = 146
	PayloadTypeLightStateHevCycleConfiguration payloadType = 147
	PayloadTypeLightGetLastHevCycleResult      payloadType = 148
	PayloadTypeLightStateLastHevCycleResult    payloadType = 149
	PayloadTypeDeviceStateUnhandled            payloadType = 223
	PayloadTypeMultiZoneSetColorZones          payloadType = 501
	PayloadTypeMultiZoneGetColorZones          payloadType = 502
	PayloadTypeMultiZoneStateZone              payloadType = 503
	PayloadTypeMultiZoneStateMultiZone         payloadType = 506
	PayloadTypeMultiZoneGetEffect              payloadType = 507
	PayloadTypeMultiZoneSetEffect              payloadType = 508
	PayloadTypeMultiZoneStateEffect            payloadType = 509
	PayloadTypeMultiZoneExtendedSetColorZones  payloadType = 510
	PayloadTypeMultiZoneExtendedGetColorZones  payloadType = 511
	PayloadTypeMultiZoneExtendedStateMultiZone payloadType = 512
	PayloadTypeTileGetDeviceChain              payloadType = 701
	PayloadTypeTileStateDeviceChain            payloadType = 702
	PayloadTypeTileSetUserPosition             payloadType = 703
	PayloadTypeTileGet64                       payloadType = 707
	PayloadTypeTileState64                     payloadType = 711
	PayloadTypeTileSet64                       payloadType = 715
	PayloadTypeTileGetEffect                   payloadType = 718
	PayloadTypeTileSetEffect                   payloadType = 719
	PayloadTypeTileStateEffect                 payloadType = 720
	PayloadTypeRelayGetPower                   payloadType = 816
	PayloadTypeRelaySetPower                   payloadType = 817
	PayloadTypeRelayStatePower                 payloadType = 818
	PayloadTypeButtonGet                       payloadType = 905
	PayloadTypeButtonSet                       payloadType = 906
	PayloadTypeButtonState                     payloadType = 907
	PayloadTypeButtonGetConfig                 payloadType = 909
	PayloadTypeButtonSetConfig                 payloadType = 910
	PayloadTypeButtonStateConfig               payloadType = 911
)

var Payloads = map[uint16]func() Payload{
	2:   func() Payload { return &DeviceGetService{} },
	3:   func() Payload { return &DeviceStateService{} },
	14:  func() Payload { return &DeviceGetHostFirmware{} },
	15:  func() Payload { return &DeviceStateHostFirmware{} },
	16:  func() Payload { return &DeviceGetWifiInfo{} },
	17:  func() Payload { return &DeviceStateWifiInfo{} },
	18:  func() Payload { return &DeviceGetWifiFirmware{} },
	19:  func() Payload { return &DeviceStateWifiFirmware{} },
	20:  func() Payload { return &DeviceGetPower{} },
	21:  func() Payload { return &DeviceSetPower{} },
	22:  func() Payload { return &DeviceStatePower{} },
	23:  func() Payload { return &DeviceGetLabel{} },
	24:  func() Payload { return &DeviceSetLabel{} },
	25:  func() Payload { return &DeviceStateLabel{} },
	32:  func() Payload { return &DeviceGetVersion{} },
	33:  func() Payload { return &DeviceStateVersion{} },
	34:  func() Payload { return &DeviceGetInfo{} },
	35:  func() Payload { return &DeviceStateInfo{} },
	38:  func() Payload { return &DeviceSetReboot{} },
	45:  func() Payload { return &DeviceAcknowledgement{} },
	48:  func() Payload { return &DeviceGetLocation{} },
	49:  func() Payload { return &DeviceSetLocation{} },
	50:  func() Payload { return &DeviceStateLocation{} },
	51:  func() Payload { return &DeviceGetGroup{} },
	52:  func() Payload { return &DeviceSetGroup{} },
	53:  func() Payload { return &DeviceStateGroup{} },
	58:  func() Payload { return &DeviceEchoRequest{} },
	59:  func() Payload { return &DeviceEchoResponse{} },
	101: func() Payload { return &LightGet{} },
	102: func() Payload { return &LightSetColor{} },
	103: func() Payload { return &LightSetWaveform{} },
	107: func() Payload { return &LightState{} },
	116: func() Payload { return &LightGetPower{} },
	117: func() Payload { return &LightSetPower{} },
	118: func() Payload { return &LightStatePower{} },
	119: func() Payload { return &LightSetWaveformOptional{} },
	120: func() Payload { return &LightGetInfrared{} },
	121: func() Payload { return &LightStateInfrared{} },
	122: func() Payload { return &LightSetInfrared{} },
	142: func() Payload { return &LightGetHevCycle{} },
	143: func() Payload { return &LightSetHevCycle{} },
	144: func() Payload { return &LightStateHevCycle{} },
	145: func() Payload { return &LightGetHevCycleConfiguration{} },
	146: func() Payload { return &LightSetHevCycleConfiguration{} },
	147: func() Payload { return &LightStateHevCycleConfiguration{} },
	148: func() Payload { return &LightGetLastHevCycleResult{} },
	149: func() Payload { return &LightStateLastHevCycleResult{} },
	223: func() Payload { return &DeviceStateUnhandled{} },
	501: func() Payload { return &MultiZoneSetColorZones{} },
	502: func() Payload { return &MultiZoneGetColorZones{} },
	503: func() Payload { return &MultiZoneStateZone{} },
	506: func() Payload { return &MultiZoneStateMultiZone{} },
	507: func() Payload { return &MultiZoneGetEffect{} },
	508: func() Payload { return &MultiZoneSetEffect{} },
	509: func() Payload { return &MultiZoneStateEffect{} },
	510: func() Payload { return &MultiZoneExtendedSetColorZones{} },
	511: func() Payload { return &MultiZoneExtendedGetColorZones{} },
	512: func() Payload { return &MultiZoneExtendedStateMultiZone{} },
	701: func() Payload { return &TileGetDeviceChain{} },
	702: func() Payload { return &TileStateDeviceChain{} },
	703: func() Payload { return &TileSetUserPosition{} },
	707: func() Payload { return &TileGet64{} },
	711: func() Payload { return &TileState64{} },
	715: func() Payload { return &TileSet64{} },
	718: func() Payload { return &TileGetEffect{} },
	719: func() Payload { return &TileSetEffect{} },
	720: func() Payload { return &TileStateEffect{} },
	816: func() Payload { return &RelayGetPower{} },
	817: func() Payload { return &RelaySetPower{} },
	818: func() Payload { return &RelayStatePower{} },
	905: func() Payload { return &ButtonGet{} },
	906: func() Payload { return &ButtonSet{} },
	907: func() Payload { return &ButtonState{} },
	909: func() Payload { return &ButtonGetConfig{} },
	910: func() Payload { return &ButtonSetConfig{} },
	911: func() Payload { return &ButtonStateConfig{} },
}
