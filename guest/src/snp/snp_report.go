package snp

import (
	"unsafe"

	"github.com/rizzza/smart/ioctl"
)

type ReportReq struct {
	UserData [64]byte
	VMPL     uint32
	Reserved [28]byte
}

type ReportResp struct {
	Data [4000]byte
}

type TCBVersion struct {
	BootLoader byte
	TEE        byte
	Reserved   [4]byte
	SNP        byte
	Microcode  byte
}

type Signature struct {
	R        [72]byte
	S        [72]byte
	Reserved [512 - 144]byte
}

type AttestationReport struct {
	Version         uint32
	GuestSVN        uint32
	Policy          uint64
	FamilyId        [16]byte
	ImageId         [16]byte
	VMPL            uint32
	SignatureAlgo   uint32
	PlatformVersion TCBVersion
	PlatformInfo    uint64
	Flags           uint32
	Reserved0       uint32
	ReportData      [64]byte
	Measurement     [48]byte
	HostData        [32]byte
	IdKeyDigest     [48]byte
	AuthorKeyDigest [48]byte
	ReportId        [32]byte
	ReportIdMA      [32]byte
	ReportedTCB     TCBVersion
	Reserved1       [24]byte
	ChipId          [64]byte
	CommitedTCB     TCBVersion
	CurrentBuild    byte
	CurrentMinor    byte
	CurrentMajor    byte
	Reserved2       byte
	CommitedBuild   byte
	CommitedMinor   byte
	CommitedMajor   byte
	Reserved3       byte
	LaunchTCB       TCBVersion
	Reserved4       [168]byte
	Signature       Signature
}

type MsgReportResp struct {
	Status            uint32
	ReportSize        uint32
	Reserved          [24]byte
	AttestationReport AttestationReport
}

var SNP_GET_REPORT_CMD = ioctl.Iowr(uintptr(SNP_GUEST_REQ_IOC_TYPE), 0x0, unsafe.Sizeof(GuestRequestIOCtl{}))
