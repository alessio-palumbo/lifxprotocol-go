// Code generated. DO NOT EDIT.
// Source: https://github.com/LIFX/public-protocol@fe8871b049893401c8c6a5fc425971ec55a2fa06
// Generated: 2025-12-16T04:29:29Z
package enums

type ButtonGesture uint16

const (
	ButtonGestureBUTTONGESTUREPRESS            ButtonGesture = 1
	ButtonGestureBUTTONGESTUREHOLD             ButtonGesture = 2
	ButtonGestureBUTTONGESTUREPRESSPRESS       ButtonGesture = 3
	ButtonGestureBUTTONGESTUREPRESSHOLD        ButtonGesture = 4
	ButtonGestureBUTTONGESTUREHOLDHOLD         ButtonGesture = 5
	ButtonGestureBUTTONGESTUREPRESSRELEASE     ButtonGesture = 6
	ButtonGestureBUTTONGESTUREHOLDRELEASE      ButtonGesture = 7
	ButtonGestureBUTTONGESTUREPRESSTOUCH       ButtonGesture = 8
	ButtonGestureBUTTONGESTUREHOLDTOUCH        ButtonGesture = 9
	ButtonGestureBUTTONGESTUREPRESSTOUCHREPEAT ButtonGesture = 10
	ButtonGestureBUTTONGESTUREHOLDTOUCHREPEAT  ButtonGesture = 11
	ButtonGestureBUTTONGESTUREPRESSFASTREPEAT  ButtonGesture = 12
)

type ButtonTargetType uint16

const (
	ButtonTargetTypeBUTTONTARGETTYPEPOWERTOGGLERELAYS         ButtonTargetType = 2
	ButtonTargetTypeBUTTONTARGETTYPEPOWERTOGGLEDEVICE         ButtonTargetType = 3
	ButtonTargetTypeBUTTONTARGETTYPEPOWERTOGGLELOCATION       ButtonTargetType = 4
	ButtonTargetTypeBUTTONTARGETTYPEPOWERTOGGLEGROUP          ButtonTargetType = 5
	ButtonTargetTypeBUTTONTARGETTYPESCENE                     ButtonTargetType = 6
	ButtonTargetTypeBUTTONTARGETTYPEPOWERTOGGLEDEVICERELAYS   ButtonTargetType = 7
	ButtonTargetTypeBUTTONTARGETTYPEBRIGHTNESSDOWNDEVICE      ButtonTargetType = 8
	ButtonTargetTypeBUTTONTARGETTYPEBRIGHTNESSDOWNGROUP       ButtonTargetType = 9
	ButtonTargetTypeBUTTONTARGETTYPEBRIGHTNESSDOWNLOCATION    ButtonTargetType = 10
	ButtonTargetTypeBUTTONTARGETTYPEBRIGHTNESSUPDEVICE        ButtonTargetType = 11
	ButtonTargetTypeBUTTONTARGETTYPEBRIGHTNESSUPGROUP         ButtonTargetType = 12
	ButtonTargetTypeBUTTONTARGETTYPEBRIGHTNESSUPLOCATION      ButtonTargetType = 13
	ButtonTargetTypeBUTTONTARGETTYPEDEMOEFFECTCYCLE           ButtonTargetType = 14
	ButtonTargetTypeBUTTONTARGETTYPEDEMOEFFECTCYCLESTOP       ButtonTargetType = 15
	ButtonTargetTypeBUTTONTARGETTYPEDEMOSUNRISESUNSET         ButtonTargetType = 16
	ButtonTargetTypeBUTTONTARGETTYPEPOWERONDEVICE             ButtonTargetType = 17
	ButtonTargetTypeBUTTONTARGETTYPEPOWERONLOCATION           ButtonTargetType = 18
	ButtonTargetTypeBUTTONTARGETTYPEPOWERONGROUP              ButtonTargetType = 19
	ButtonTargetTypeBUTTONTARGETTYPEPOWERONRELAYS             ButtonTargetType = 20
	ButtonTargetTypeBUTTONTARGETTYPEPOWEROFFDEVICE            ButtonTargetType = 21
	ButtonTargetTypeBUTTONTARGETTYPEPOWEROFFLOCATION          ButtonTargetType = 22
	ButtonTargetTypeBUTTONTARGETTYPEPOWEROFFGROUP             ButtonTargetType = 23
	ButtonTargetTypeBUTTONTARGETTYPEPOWEROFFRELAYS            ButtonTargetType = 24
	ButtonTargetTypeBUTTONTARGETTYPEPOWERTOGGLELOCALDEVICE    ButtonTargetType = 28
	ButtonTargetTypeBUTTONTARGETTYPEBRIGHTNESSDOWNLOCALDEVICE ButtonTargetType = 29
	ButtonTargetTypeBUTTONTARGETTYPEBRIGHTNESSUPLOCALDEVICE   ButtonTargetType = 30
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
