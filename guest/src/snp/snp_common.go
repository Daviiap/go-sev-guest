package snp

type TCBVersion struct {
	BootLoader byte
	TEE        byte
	Reserved   [4]byte
	SNP        byte
	Microcode  byte
}
