package commands

import (
	"bytes"
	"encoding/binary"
	"unsafe"

	snp "sev-guest/src/snp"

	"github.com/rizzza/smart/ioctl"
)

type KeyDerivationOptions struct {
	KeySel           uint8
	RootKeySel       uint8
	GuestFieldSelect struct {
		GuestPolicy bool
		ImageID     bool
		FamilyID    bool
		Measurement bool
		GuestSVN    bool
		TCBVersion  bool
	}
	VMPL       uint32
	GuestSVN   uint32
	TCBVersion snp.TCBVersion
}

const ROOT_KEY_SEL_VCEK = 0b000
const ROOT_KEY_SEL_VMRK = 0b001

const KEY_SEL_VLEK_OR_VCEK = 0b000
const KEY_SEL_VCEK = 0b010
const KEY_SEL_VLEK = 0b100

const GUEST_FIELD_SELECT_GUEST_POLICY = 0b000001
const GUEST_FIELD_SELECT_IMAGE_ID = 0b000010
const GUEST_FIELD_SELECT_FAMILY_ID = 0b000100
const GUEST_FIELD_SELECT_MEASUREMENT = 0b001000
const GUEST_FIELD_SELECT_GUEST_SVN = 0b010000
const GUEST_FIELD_SELECT_TCB_VERSION = 0b100000

func DeriveKey(options KeyDerivationOptions) (snp.MsgDeriveKeyResp, error) {
	var req snp.DeriveKeyReq
	var resp snp.DeriveKeyResp
	var guestReq snp.GuestRequestIOCtl
	var reportResp snp.MsgDeriveKeyResp

	guestReq.MSGVersion = 0x01
	guestReq.ReqData = uint64(uintptr(unsafe.Pointer(&req)))
	guestReq.RespData = uint64(uintptr(unsafe.Pointer(&resp)))

	var SNP_DERIVE_KEY_CMD = ioctl.Iowr(uintptr(snp.SNP_GUEST_REQ_IOC_TYPE), 0x01, unsafe.Sizeof(snp.GuestRequestIOCtl{}))

	snp.SNPIOCtl(&guestReq, SNP_DERIVE_KEY_CMD)

	binary.Read(bytes.NewBuffer(resp.Data[:]), binary.LittleEndian, &reportResp)

	return reportResp, nil
}
