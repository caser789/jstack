package udp

import "github.com/caser789/jstack/tcpip/stack"
import "github.com/caser789/jstack/tcpip/buffer"
import "github.com/caser789/jstack/tcpip/header"
import "github.com/caser789/jstack/tcpip"

type endpoint struct {
	stack     *stack.Stack
	bindNICID tcpip.NICID
	netProto  tcpip.NetworkProtocolNumber
	bindAddr  tcpip.Address
	id        stack.TransportEndpointID

	rcvList    udpPacketList
	rcvBufSize int
}

func newEndpoint(netProto tcpip.NetworkProtocolNumber) *endpoint {
	return &endpoint{
		netProto: netProto,
	}
}

func (e *endpoint) Close() {
}

func (e *endpoint) Bind(addr tcpip.FullAddress) error {
	e.bindNICID = addr.NIC
	e.bindAddr = addr.Addr
	return nil
}

func (*endpoint) Listen(int) error {
	return tcpip.ErrNotSupported
}

func (*endpoint) Accept() error {
	return tcpip.ErrNotSupported
}

func (e *endpoint) Write(v buffer.View, to *tcpip.FullAddress) (uintptr, error) {
	route, _ := e.stack.FindRoute(e.bindNICID, e.bindAddr, to.Addr, e.netProto)
	defer route.Release()

	sendUDP(&route, v, e.id.LocalPort, to.Port)
	return uintptr(len(v)), nil
}

func sendUDP(r *stack.Route, data buffer.View, localPort, remotePort uint16) error {
	hdr := buffer.NewPrependable(header.UDPMinimumSize + int(r.MaxHeaderLength()))
	udp := header.UDP(hdr.Prepend(header.UDPMinimumSize))
	length := uint16(hdr.UsedLength())
	xsum := r.PseudoHeaderChecksum(ProtocolNumber)
	if data != nil {
		length += uint16(len(data))
		xsum = header.Checksum(data, xsum)
	}
	udp.Encode(&header.UDPFields{
		SrcPort: localPort,
		DstPort: remotePort,
		Length:  length,
	})
	udp.SetChecksum(^udp.CalculateChecksum(xsum, length))
	return r.WritePacket(&hdr, data, ProtocolNumber)
}

func (e *endpoint) RecvMsg(addr *tcpip.FullAddress) (buffer.View, error) {
	v, err := e.Read(addr)
	return v, err
}

func (e *endpoint) Read(addr *tcpip.FullAddress) (buffer.View, error) {
	if e.rcvList.Empty() {
		err := tcpip.ErrWouldBlock
		return buffer.View{}, err
	}

	p := e.rcvList.Front()
	e.rcvList.Remove(p)
	e.rcvBufSize -= len(p.view)

	if addr != nil {
		*addr = p.senderAddress
	}

	return p.view, nil
}
