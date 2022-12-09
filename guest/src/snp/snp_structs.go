package snp

import (
	"os"
	"unsafe"

	"github.com/rizzza/smart/ioctl"
)

type SnpGuestRequestIOCtl struct {
	MSGVersion byte
	ReqData    uint64
	RespData   uint64
	FWErr      uint64
}

func getSNPDevFD() (uintptr, error) {
	file, err := os.Open("/dev/sev-guest")

	if err != nil {
		return 0, err
	}

	defer file.Close()

	fd := file.Fd()

	return fd, nil
}

func SNPIOCtl(guestReq *SnpGuestRequestIOCtl) error {
	fd, err := getSNPDevFD()

	if err != nil {
		return err
	}

	const SNP_GUEST_REQ_IOC_TYPE = 'S'
	var SNP_GET_REPORT = ioctl.Iowr(uintptr(SNP_GUEST_REQ_IOC_TYPE), 0x0, unsafe.Sizeof(SnpGuestRequestIOCtl{}))

	return ioctl.Ioctl(fd, SNP_GET_REPORT, uintptr(unsafe.Pointer(guestReq)))
}
