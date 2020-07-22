package stack

import "github.com/caser789/jstack/tcpip"

func New(network []string, transport []string) IStack {
    s := &Stack{
		transportProtocols: make(map[tcpip.TransportProtocolNumber]ITransportProtocol),
		networkProtocols:   make(map[tcpip.NetworkProtocolNumber]NetworkProtocol),
    }

	for _, name := range network {
		netProto, ok := networkProtocols[name]
		if !ok {
			continue
		}

		s.networkProtocols[netProto.Number()] = netProto
	}


    for _, name := range transport {
        proto, ok := transportProtocols[name]
        if !ok {
            continue
        }

        s.transportProtocols[proto.Number()] = proto
    }
    return s
}


// global
var linkEndpoints = make(map[tcpip.LinkEndpointID]ILinkEndpoint)
func FindLinkEndpoint(id tcpip.LinkEndpointID) ILinkEndpoint {
	return linkEndpoints[id]
}
