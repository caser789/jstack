package udp

import "github.com/caser789/jstack/tcpip/stack"
import "github.com/caser789/jstack/tcpip/header"
import "github.com/caser789/jstack/tcpip"

const ProtocolNumber = header.UDPProtocolNumber
const ProtocolName = "udp"

type protocol struct{}

func (*protocol) Number() tcpip.TransportProtocolNumber {
	return ProtocolNumber
}

func (*protocol) NewEndpoint(netProto tcpip.NetworkProtocolNumber) (stack.IEndpoint, error) {
	return newEndpoint(netProto), nil
}

func init() {
	stack.RegisterTransportProtocol(ProtocolName, &protocol{})
}
