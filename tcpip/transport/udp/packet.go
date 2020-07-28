package udp

import "github.com/caser789/jstack/tcpip/buffer"
import "github.com/caser789/jstack/tcpip"

type udpPacketEntry struct {
	next *udpPacket
	prev *udpPacket
}

func (e *udpPacketEntry) Next() *udpPacket {
	return e.next
}

func (e *udpPacketEntry) Prev() *udpPacket {
	return e.prev
}

func (e *udpPacketEntry) SetNext(entry *udpPacket) {
	e.next = entry
}

func (e *udpPacketEntry) SetPrev(entry *udpPacket) {
	e.prev = entry
}

type udpPacket struct {
	udpPacketEntry
	senderAddress tcpip.FullAddress
	view          buffer.View
}
