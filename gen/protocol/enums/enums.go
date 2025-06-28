// Code generated. DO NOT EDIT.
// Source: https://github.com/LIFX/public-protocol@f061a60a2e9305e63ca75cfa53ad78d276de0e0a
// Generated: 2025-06-28T09:47:36Z
package enums

type ButtonGesture uint16

const (
	ButtonGestureBUTTONGESTUREPRESS      ButtonGesture = 1
	ButtonGestureBUTTONGESTUREHOLD       ButtonGesture = 2
	ButtonGestureBUTTONGESTUREPRESSPRESS ButtonGesture = 3
	ButtonGestureBUTTONGESTUREPRESSHOLD  ButtonGesture = 4
	ButtonGestureBUTTONGESTUREHOLDHOLD   ButtonGesture = 5
)

type ButtonTargetType uint16

const (
	ButtonTargetTypeBUTTONTARGETTYPERELAYS       ButtonTargetType = 2
	ButtonTargetTypeBUTTONTARGETTYPEDEVICE       ButtonTargetType = 3
	ButtonTargetTypeBUTTONTARGETTYPELOCATION     ButtonTargetType = 4
	ButtonTargetTypeBUTTONTARGETTYPEGROUP        ButtonTargetType = 5
	ButtonTargetTypeBUTTONTARGETTYPESCENE        ButtonTargetType = 6
	ButtonTargetTypeBUTTONTARGETTYPEDEVICERELAYS ButtonTargetType = 7
)

type DeviceService uint8

const (
	DeviceServiceDEVICESERVICEUDP DeviceService = 1
)

type LightWaveform uint8

const (
	LightWaveformLIGHTWAVEFORMSAW      LightWaveform = 0
	LightWaveformLIGHTWAVEFORMSINE     LightWaveform = 1
	LightWaveformLIGHTWAVEFORMHALFSINE LightWaveform = 2
	LightWaveformLIGHTWAVEFORMTRIANGLE LightWaveform = 3
	LightWaveformLIGHTWAVEFORMPULSE    LightWaveform = 4
)

type LightLastHevCycleResult uint8

const (
	LightLastHevCycleResultLIGHTLASTHEVCYCLERESULTSUCCESS              LightLastHevCycleResult = 0
	LightLastHevCycleResultLIGHTLASTHEVCYCLERESULTBUSY                 LightLastHevCycleResult = 1
	LightLastHevCycleResultLIGHTLASTHEVCYCLERESULTINTERRUPTEDBYRESET   LightLastHevCycleResult = 2
	LightLastHevCycleResultLIGHTLASTHEVCYCLERESULTINTERRUPTEDBYHOMEKIT LightLastHevCycleResult = 3
	LightLastHevCycleResultLIGHTLASTHEVCYCLERESULTINTERRUPTEDBYLAN     LightLastHevCycleResult = 4
	LightLastHevCycleResultLIGHTLASTHEVCYCLERESULTINTERRUPTEDBYCLOUD   LightLastHevCycleResult = 5
	LightLastHevCycleResultLIGHTLASTHEVCYCLERESULTNONE                 LightLastHevCycleResult = 255
)

type MultiZoneApplicationRequest uint8

const (
	MultiZoneApplicationRequestMULTIZONEAPPLICATIONREQUESTNOAPPLY   MultiZoneApplicationRequest = 0
	MultiZoneApplicationRequestMULTIZONEAPPLICATIONREQUESTAPPLY     MultiZoneApplicationRequest = 1
	MultiZoneApplicationRequestMULTIZONEAPPLICATIONREQUESTAPPLYONLY MultiZoneApplicationRequest = 2
)

type MultiZoneEffectType uint8

const (
	MultiZoneEffectTypeMULTIZONEEFFECTTYPEOFF  MultiZoneEffectType = 0
	MultiZoneEffectTypeMULTIZONEEFFECTTYPEMOVE MultiZoneEffectType = 1
)

type MultiZoneExtendedApplicationRequest uint8

const (
	MultiZoneExtendedApplicationRequestMULTIZONEEXTENDEDAPPLICATIONREQUESTNOAPPLY   MultiZoneExtendedApplicationRequest = 0
	MultiZoneExtendedApplicationRequestMULTIZONEEXTENDEDAPPLICATIONREQUESTAPPLY     MultiZoneExtendedApplicationRequest = 1
	MultiZoneExtendedApplicationRequestMULTIZONEEXTENDEDAPPLICATIONREQUESTAPPLYONLY MultiZoneExtendedApplicationRequest = 2
)

type TileEffectSkyPalette uint8

const (
	TileEffectSkyPaletteTILEEFFECTSKYPALETTECLOUDSSKY  TileEffectSkyPalette = 0
	TileEffectSkyPaletteTILEEFFECTSKYPALETTENIGHTSKY   TileEffectSkyPalette = 1
	TileEffectSkyPaletteTILEEFFECTSKYPALETTEDAWNSKY    TileEffectSkyPalette = 2
	TileEffectSkyPaletteTILEEFFECTSKYPALETTEDAWNSUN    TileEffectSkyPalette = 3
	TileEffectSkyPaletteTILEEFFECTSKYPALETTEFULLSUN    TileEffectSkyPalette = 4
	TileEffectSkyPaletteTILEEFFECTSKYPALETTEFINALSUN   TileEffectSkyPalette = 5
	TileEffectSkyPaletteTILEEFFECTSKYPALETTENUMCOLOURS TileEffectSkyPalette = 6
)

type TileEffectSkyType uint8

const (
	TileEffectSkyTypeTILEEFFECTSKYTYPESUNRISE TileEffectSkyType = 0
	TileEffectSkyTypeTILEEFFECTSKYTYPESUNSET  TileEffectSkyType = 1
	TileEffectSkyTypeTILEEFFECTSKYTYPECLOUDS  TileEffectSkyType = 2
)

type TileEffectType uint8

const (
	TileEffectTypeTILEEFFECTTYPEOFF   TileEffectType = 0
	TileEffectTypeTILEEFFECTTYPEMORPH TileEffectType = 2
	TileEffectTypeTILEEFFECTTYPEFLAME TileEffectType = 3
	TileEffectTypeTILEEFFECTTYPESKY   TileEffectType = 5
)
