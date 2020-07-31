package stack

import "github.com/caser789/jstack/tcpip"
import "github.com/caser789/jstack/tcpip/buffer"

type INetworkDispatcher interface{}
type ITransportDispatcher interface{}

type IEndpoint interface {
	Close()
	Bind(address tcpip.FullAddress) error
	RecvMsg(*tcpip.FullAddress) (buffer.View, error)

	Write(buffer.View, *FullAddress) (uintptr, error)
}

type ILinkEndpoint interface {
	Attach(dispatcher INetworkDispatcher)
}
type INetworkEndpoint interface {
	MTU() uint32
	MaxHeaderLength() uint16
	ID() *NetworkEndpointID
	NICID() tcpip.NICID

	WritePacket(r *Route, hdr *buffer.Prependable, payload buffer.View, protocol tcpip.TransportProtocolNumber) error
	HandlePacket(r *Route, v buffer.View)
}

type IStack interface {
	NewEndpoint(transport tcpip.TransportProtocolNumber, network tcpip.NetworkProtocolNumber) (IEndpoint, error)
}

type ITransportProtocol interface {
	NewEndpoint(netProto tcpip.NetworkProtocolNumber) (IEndpoint, error)
	Number() tcpip.TransportProtocolNumber
}
type INetworkProtocol interface {
	NewEndpoint(nicid tcpip.NICID, addr tcpip.Address, dispatcher ITransportDispatcher, sender ILinkEndpoint) (INetworkEndpoint, error)
	Number() tcpip.NetworkProtocolNumber
}

type NetworkEndpointID struct {
	LocalAddress tcpip.Address
}
type TransportEndpointID struct {
	LocalPort    uint16
	LocalAddress tcpip.Address

	RemotePort    uint16
	RemoteAddress tcpip.Address
}
