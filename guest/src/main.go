package main

func main() {
	report := AttestationReport{}
	GetReport([64]byte{}, &report)
	PrintAttestationReport(&report)
}
