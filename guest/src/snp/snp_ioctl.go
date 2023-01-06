package snp

import (
	"os"
	"unsafe"

	"github.com/rizzza/smart/ioctl"
)

type GuestRequestIOCtl struct {
	MSGVersion byte
	ReqData    uint64
	RespData   uint64
	FWErr      uint64
}

const SNP_GUEST_REQ_IOC_TYPE = 'S'

func SNPIOCtl(guestReq *GuestRequestIOCtl, cmd uintptr) error {
	file, err := os.Open("/dev/sev-guest")

	if err != nil {
		return err
	}

	defer file.Close()

	fd := file.Fd()

	return ioctl.Ioctl(fd, cmd, uintptr(unsafe.Pointer(guestReq)))
}
