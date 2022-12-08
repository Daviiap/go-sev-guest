package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"unsafe"

	"github.com/rizzza/smart/ioctl"
)

type SnpReportReq struct {
	UserData [64]byte
	VMPL     uint32
	Reserved [28]byte
}

type SnpReportResp struct {
	Data [4000]byte
}

type SnpGuestRequestIOCtl struct {
	MSGVersion byte
	ReqData    uint64
	RespData   uint64
	FWErr      uint64
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

func GetReport(data [64]byte, report *AttestationReport) ([]byte, error) {
	var req SnpReportReq
	var resp SnpReportResp
	var guestReq SnpGuestRequestIOCtl
	var reportResp MsgReportResp

	req.UserData = data

	guestReq.MSGVersion = 0x01
	guestReq.ReqData = uint64(uintptr(unsafe.Pointer(&req)))
	guestReq.RespData = uint64(uintptr(unsafe.Pointer(&resp)))

	file, err := os.Open("/dev/sev-guest")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	fd := file.Fd()

	const SNP_GUEST_REQ_IOC_TYPE = 'S'
	var SNP_GET_REPORT = ioctl.Iowr(uintptr(SNP_GUEST_REQ_IOC_TYPE), 0x0, unsafe.Sizeof(SnpGuestRequestIOCtl{}))

	err = ioctl.Ioctl(fd, SNP_GET_REPORT, uintptr(unsafe.Pointer(&guestReq)))

	if err != nil {
		return nil, err
	}

	reportBin := resp.Data[32 : 1184+32]

	binary.Read(bytes.NewBuffer(resp.Data[:]), binary.LittleEndian, &reportResp)
	binary.Read(bytes.NewBuffer(reportBin), binary.LittleEndian, report)

	if reportResp.Status != 0 {
		return nil, fmt.Errorf("error: status: %d", reportResp.Status)
	}

	reportSize := unsafe.Sizeof(*report)
	if reportResp.ReportSize > uint32(reportSize) {
		return nil, fmt.Errorf("error: received report size: %d, expected %d", reportResp.ReportSize, reportSize)
	}

	return reportBin, nil
}

func PrintByteArray(array []byte) string {
	str := ""

	for i := 0; i < len(array); i++ {
		value := array[i]
		str += fmt.Sprintf("%02x", value)
	}

	return str
}

func PrintAttestationReport(report *AttestationReport) {
	fmt.Print("Version: ")
	fmt.Println(report.Version)
	fmt.Print("GuestSVN: ")
	fmt.Println(report.GuestSVN)
	fmt.Print("Policy: ")
	fmt.Printf("0x%x\n", report.Policy)
	fmt.Print("FamilyId: ")
	fmt.Println(PrintByteArray(report.FamilyId[:]))
	fmt.Print("ImageId: ")
	fmt.Println(PrintByteArray(report.ImageId[:]))
	fmt.Print("VMPL: ")
	fmt.Println(report.VMPL)
	fmt.Print("SignatureAlgo: ")
	fmt.Println(report.SignatureAlgo)
	fmt.Print("CurrentTCB: ")
	fmt.Print("    BootLoader: ")
	fmt.Println(report.PlatformVersion.BootLoader)
	fmt.Print("    Microcode: ")
	fmt.Println(report.PlatformVersion.Microcode)
	fmt.Print("    SNP: ")
	fmt.Println(report.PlatformVersion.SNP)
	fmt.Print("    TEE: ")
	fmt.Println(report.PlatformVersion.TEE)
	fmt.Print("PlatformInfo: ")
	fmt.Println(report.PlatformInfo)
	fmt.Print("AuthorKeyDigest: ")
	fmt.Println(PrintByteArray((report.AuthorKeyDigest[:])))
	fmt.Print("ReportData: ")
	fmt.Println(PrintByteArray((report.ReportData[:])))
	fmt.Print("Measurement: ")
	fmt.Println(PrintByteArray((report.Measurement[:])))
	fmt.Print("HostData: ")
	fmt.Println(PrintByteArray((report.HostData[:])))
	fmt.Print("IdKeyDigest: ")
	fmt.Println(PrintByteArray((report.IdKeyDigest[:])))
	fmt.Print("AuthorKeyDigest: ")
	fmt.Println(PrintByteArray((report.AuthorKeyDigest[:])))
	fmt.Print("ReportId: ")
	fmt.Println(PrintByteArray((report.ReportId[:])))
	fmt.Print("ReportIdMA: ")
	fmt.Println(PrintByteArray((report.ReportIdMA[:])))
	fmt.Print("ReportedTCB: ")
	fmt.Print("    BootLoader: ")
	fmt.Println(report.ReportedTCB.BootLoader)
	fmt.Print("    Microcode: ")
	fmt.Println(report.ReportedTCB.Microcode)
	fmt.Print("    SNP: ")
	fmt.Println(report.ReportedTCB.SNP)
	fmt.Print("    TEE: ")
	fmt.Println(report.ReportedTCB.TEE)
	fmt.Print("ChipId: ")
	fmt.Println(PrintByteArray((report.ChipId[:])))
	fmt.Print("CurrentBuild: ")
	fmt.Println(report.CurrentBuild)
	fmt.Print("CurrentMinor: ")
	fmt.Println(report.CurrentMinor)
	fmt.Print("CurrentMajor: ")
	fmt.Println(report.CurrentMajor)
	fmt.Print("CommitedTCB: ")
	fmt.Print("    BootLoader: ")
	fmt.Println(report.CommitedTCB.BootLoader)
	fmt.Print("    Microcode: ")
	fmt.Println(report.CommitedTCB.Microcode)
	fmt.Print("    SNP: ")
	fmt.Println(report.CommitedTCB.SNP)
	fmt.Print("    TEE: ")
	fmt.Println(report.CommitedTCB.TEE)
	fmt.Print("CommitedBuild: ")
	fmt.Println(report.CommitedBuild)
	fmt.Print("CommitedMinor: ")
	fmt.Println(report.CommitedMinor)
	fmt.Print("CommitedMajor: ")
	fmt.Println(report.CommitedMajor)
	fmt.Print("LaunchTCB: ")
	fmt.Print("    BootLoader: ")
	fmt.Println(report.LaunchTCB.BootLoader)
	fmt.Print("    Microcode: ")
	fmt.Println(report.LaunchTCB.Microcode)
	fmt.Print("    SNP: ")
	fmt.Println(report.LaunchTCB.SNP)
	fmt.Print("    TEE: ")
	fmt.Println(report.LaunchTCB.TEE)
}
