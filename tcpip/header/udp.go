package header

import "github.com/caser789/jstack/tcpip"

const UDPProtocolNumber tcpip.TransportProtocolNumber = 17
const UDPMinimumSize = 8

type UDP []byte

type UDPFields struct {
	SrcPort uint16
	DstPort uint16
	Length uint16
	Checksum uint16
}


func (b UDP) Encode(u *UDPFields) {
}

func (b UDP) SetChecksum(checksum uint16) {
}
func (b UDP) CalculateChecksum(partialChecksum uint16, totalLen uint16) uint16 {
    return 0
}
