package stack

import "github.com/caser789/jstack/tcpip"
import "github.com/caser789/jstack/tcpip/header"
import "github.com/caser789/jstack/tcpip/buffer"

type Route struct {
	RemoteAddress tcpip.Address
	LocalAddress  tcpip.Address
	NextHop       tcpip.Address
	NetProto      tcpip.NetworkProtocolNumber
	ref           INetworkEndpoint
}

func makeRoute(netProto tcpip.NetworkProtocolNumber, localAddr, remoteAddr tcpip.Address, ref INetworkEndpoint) Route {
	return Route{
		NetProto:      netProto,
		LocalAddress:  localAddr,
		RemoteAddress: remoteAddr,
		ref:           ref,
	}
}

func (r *Route) Release() {}

func (r *Route) MaxHeaderLength() uint16 {
	return r.ref.MaxHeaderLength()
}

func (r *Route) PseudoHeaderChecksum(protocol tcpip.TransportProtocolNumber) uint16 {
	return header.PseudoHeaderChecksum(protocol, r.LocalAddress, r.RemoteAddress)
}

func (r *Route) WritePacket(hdr *buffer.Prependable, payload buffer.View, protocol tcpip.TransportProtocolNumber) error {
	return r.ref.WritePacket(r, hdr, payload, protocol)
}
