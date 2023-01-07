package snp

import (
	"unsafe"

	"github.com/rizzza/smart/ioctl"
)

type DeriveKeyReq struct {
	Selection        uint32
	Reserved         uint32
	GuestFieldSelect uint64
	VMPL             uint32
	GuestSVN         uint32
	TCBVersion       TCBVersion
}

type MsgDeriveKeyResp struct {
	Status     uint32
	Reserved   [28]byte
	DerivedKey [32]byte
}

type DeriveKeyResp struct {
	Data [64]byte
}

var SNP_DERIVE_KEY_CMD = ioctl.Iowr(uintptr(SNP_GUEST_REQ_IOC_TYPE), 0x01, unsafe.Sizeof(GuestRequestIOCtl{}))
