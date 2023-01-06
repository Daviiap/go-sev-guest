package commands

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"sev-guest/src/snp"
)

const POLICY_SINGLE_SOCKET_SHIFT = 20
const POLICY_DEBUG_SHIFT = 19
const POLICY_MIGRATE_MA_SHIFT = 18
const POLICY_SMT_SHIFT = 16
const POLICY_ABI_MAJOR_SHIFT = 8
const POLICY_ABI_MINOR_SHIFT = 0

const POLICY_SINGLE_SOCKET_MASK = (uint64(1) << (POLICY_SINGLE_SOCKET_SHIFT))
const POLICY_DEBUG_MASK = (uint64(1) << (POLICY_DEBUG_SHIFT))
const POLICY_MIGRATE_MA_MASK = (uint64(1) << (POLICY_MIGRATE_MA_SHIFT))
const POLICY_SMT_MASK = (uint64(1) << (POLICY_SMT_SHIFT))
const POLICY_ABI_MAJOR_MASK = (uint64(0xff) << (POLICY_ABI_MAJOR_SHIFT))
const POLICY_ABI_MINOR_MASK = (uint64(0xff) << (POLICY_ABI_MINOR_SHIFT))

func ReadReport(reportPath string, report *snp.AttestationReport) error {
	reportBin, err := os.ReadFile(reportPath)

	if err != nil {
		return err
	}

	binary.Read(bytes.NewBuffer(reportBin), binary.LittleEndian, report)

	return nil
}

func PrintByteArray(array []byte) string {
	str := ""

	for i := 0; i < len(array); i++ {
		value := array[i]
		str += fmt.Sprintf("%02x", value)
	}

	return str
}

func PrintAttestationReport(report *snp.AttestationReport) {
	fmt.Print("Version: ")
	fmt.Println(report.Version)
	fmt.Print("GuestSVN: ")
	fmt.Println(report.GuestSVN)
	fmt.Print("Policy: ")
	fmt.Printf("0x%x\n", report.Policy)
	fmt.Print("    ABIMinor: ")
	fmt.Println(report.Policy & POLICY_ABI_MINOR_MASK >> POLICY_ABI_MINOR_SHIFT)
	fmt.Print("    ABIMajor: ")
	fmt.Println((report.Policy & POLICY_ABI_MAJOR_MASK) >> POLICY_ABI_MAJOR_SHIFT)
	fmt.Print("    SMT: ")
	fmt.Println((report.Policy & POLICY_SMT_MASK) != 0)
	fmt.Print("    MigrateMA: ")
	fmt.Println((report.Policy & POLICY_MIGRATE_MA_MASK) != 0)
	fmt.Print("    Debug: ")
	fmt.Println((report.Policy & POLICY_DEBUG_MASK) != 0)
	fmt.Print("    SingleSocket: ")
	fmt.Println((report.Policy & POLICY_SINGLE_SOCKET_MASK) != 0)
	fmt.Print("FamilyId: ")
	fmt.Println(PrintByteArray(report.FamilyId[:]))
	fmt.Print("ImageId: ")
	fmt.Println(PrintByteArray(report.ImageId[:]))
	fmt.Print("VMPL: ")
	fmt.Println(report.VMPL)
	fmt.Print("SignatureAlgo: ")
	fmt.Println(report.SignatureAlgo)
	fmt.Println("CurrentTCB: ")
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
	fmt.Println("ReportedTCB: ")
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
	fmt.Println("CommitedTCB: ")
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
	fmt.Println("LaunchTCB: ")
	fmt.Print("    BootLoader: ")
	fmt.Println(report.LaunchTCB.BootLoader)
	fmt.Print("    Microcode: ")
	fmt.Println(report.LaunchTCB.Microcode)
	fmt.Print("    SNP: ")
	fmt.Println(report.LaunchTCB.SNP)
	fmt.Print("    TEE: ")
	fmt.Println(report.LaunchTCB.TEE)
}

type ReadReportOptions struct {
	Filename string
}

func ReadReportCommand(options ReadReportOptions) {
	report := snp.AttestationReport{}
	err := ReadReport(options.Filename, &report)

	if err != nil {
		log.Fatal(err)
	}

	PrintAttestationReport(&report)
}
