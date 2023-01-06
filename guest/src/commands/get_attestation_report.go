package commands

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"os"
	"unsafe"

	snp "sev-guest/src/snp"
)

func GetReport(data [64]byte) ([]byte, error) {
	var req snp.SnpReportReq
	var resp snp.SnpReportResp
	var guestReq snp.SnpGuestRequestIOCtl
	var reportResp snp.MsgReportResp

	req.UserData = data

	guestReq.MSGVersion = 0x01
	guestReq.ReqData = uint64(uintptr(unsafe.Pointer(&req)))
	guestReq.RespData = uint64(uintptr(unsafe.Pointer(&resp)))

	err := snp.SNPIOCtl(&guestReq, snp.SNP_GET_REPORT_CMD)

	if err != nil {
		return nil, err
	}

	reportBin := resp.Data[32 : 1184+32]

	binary.Read(bytes.NewBuffer(resp.Data[:]), binary.LittleEndian, &reportResp)

	if reportResp.Status != 0 {
		return nil, fmt.Errorf("error: status: %d", reportResp.Status)
	}

	reportSize := unsafe.Sizeof(snp.AttestationReport{})
	if reportResp.ReportSize != uint32(reportSize) {
		return nil, fmt.Errorf("error: received report size: %d, expected %d", reportResp.ReportSize, reportSize)
	}

	return reportBin, nil
}

func WriteAttestationReport(report *[]byte, fileName string) error {
	reportSize := int(unsafe.Sizeof(snp.AttestationReport{}))
	if len(*report) != reportSize {
		return fmt.Errorf("error: received report size: %d, expected %d", len(*report), reportSize)
	}

	if fileName == "" {
		return errors.New("error: must provide a valid filename (size > 0)")
	}

	return os.WriteFile(fileName, *report, 0x04)
}

type GetReportOptions struct {
	Filename     string
	DataFileName string
}

func GetReportCommand(options GetReportOptions) {
	reportBin, err := GetReport([64]byte{})

	if err != nil {
		log.Fatal(err)
	}

	WriteAttestationReport(&reportBin, options.Filename)
}
