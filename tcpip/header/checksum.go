package header

import "github.com/caser789/jstack/tcpip"

func PseudoHeaderChecksum(protocol tcpip.TransportProtocolNumber, srcAddr tcpip.Address, dstAddr tcpip.Address) uint16 {
	xsum := Checksum([]byte(srcAddr), 0)
	xsum = Checksum([]byte(dstAddr), xsum)
	return Checksum([]byte{0, uint8(protocol)}, xsum)
}

func Checksum(buf []byte, initial uint16) uint16 {
	v := uint32(initial)

	l := len(buf)
	if l&1 != 0 {
		l--
		v += uint32(buf[l]) << 8
	}

	for i := 0; i < l; i += 2 {
		v += (uint32(buf[i]) << 8) + uint32(buf[i+1])
	}

	return ChecksumCombine(uint16(v), uint16(v>>16))
}

func ChecksumCombine(a, b uint16) uint16 {
	v := uint32(a) + uint32(b)
	return uint16(v + v>>16)
}
