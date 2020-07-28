package header

import "github.com/caser789/jstack/tcpip"
import "encoding/binary"

const UDPProtocolNumber tcpip.TransportProtocolNumber = 17
const UDPMinimumSize = 8

const (
	udpSrcPort  = 0
	udpDstPort  = 2
	udpLength   = 4
	udpChecksum = 6
)

type UDP []byte

type UDPFields struct {
	SrcPort  uint16
	DstPort  uint16
	Length   uint16
	Checksum uint16
}

func (b UDP) Encode(u *UDPFields) {
	binary.BigEndian.PutUint16(b[udpSrcPort:], u.SrcPort)
	binary.BigEndian.PutUint16(b[udpDstPort:], u.DstPort)
	binary.BigEndian.PutUint16(b[udpLength:], u.Length)
	binary.BigEndian.PutUint16(b[udpChecksum:], u.Checksum)
}

func (b UDP) CalculateChecksum(partialChecksum uint16, totalLen uint16) uint16 {
	tmp := make([]byte, 2)
	binary.BigEndian.PutUint16(tmp, totalLen)
	checksum := Checksum(tmp, partialChecksum)

	return Checksum(b[:UDPMinimumSize], checksum)
}

func (b UDP) SourcePort() uint16 {
	return binary.BigEndian.Uint16(b[udpSrcPort:])
}

func (b UDP) SetSourcePort(port uint16) {
	binary.BigEndian.PutUint16(b[udpSrcPort:], port)
}

func (b UDP) DestinationPort() uint16 {
	return binary.BigEndian.Uint16(b[udpDstPort:])
}

func (b UDP) SetDestinationPort(port uint16) {
	binary.BigEndian.PutUint16(b[udpDstPort:], port)
}

func (b UDP) Length() uint16 {
	return binary.BigEndian.Uint16(b[udpLength:])
}

func (b UDP) Checksum() uint16 {
	return binary.BigEndian.Uint16(b[udpChecksum:])
}

func (b UDP) SetChecksum(checksum uint16) {
	binary.BigEndian.PutUint16(b[udpChecksum:], checksum)
}

func (b UDP) Payload() []byte {
	return b[UDPMinimumSize:]
}
