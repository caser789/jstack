package stack

import "github.com/caser789/jstack/tcpip"

type Stack struct {
    transportProtocols map[tcpip.TransportProtocolNumber]ITransportProtocol
	routeTable []tcpip.Route
	nics map[tcpip.NICID]*NIC
}

func (s *Stack) NewEndpoint(transport tcpip.TransportProtocolNumber, network tcpip.NetworkProtocolNumber) (IEndpoint, error) {
	proto, ok := s.transportProtocols[transport]
	if !ok {
		return nil, ErrUnknownProtocol
	}

	return proto.NewEndpoint(network)
}

func (s *Stack) FindRoute(id tcpip.NICID, localAddr, remoteAddr tcpip.Address, netProto tcpip.NetworkProtocolNumber) (Route, error) {
	for i := range s.routeTable {
        if id != 0 && id != s.routeTable[i].NIC {
            continue
        }

        if !s.routeTable[i].Match(remoteAddr) {
            continue
        }

        nic := s.nics[s.routeTable[i].NIC]
        if nic == nil {
            continue
        }

        ep := nic.findEndpoint(localAddr)
        if ep == nil {
            continue
        }

        return makeRoute(netProto, ep.ID().LocalAddress, remoteAddr, ep), nil
    }

    return Route{}, tcpip.ErrNoRoute
}
