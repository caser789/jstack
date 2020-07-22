package stack

import "github.com/caser789/jstack/tcpip"

type Stack struct {
    // need to register the protocols before using NewEndpoint
    transportProtocols map[tcpip.TransportProtocolNumber]ITransportProtocol
	networkProtocols   map[tcpip.NetworkProtocolNumber]INetworkProtocol

    // need to register before FindRoute
	routeTable []tcpip.Route
	nics map[tcpip.NICID]*NIC
}

// new TransportEndpoint
func (s *Stack) NewEndpoint(transport tcpip.TransportProtocolNumber, network tcpip.NetworkProtocolNumber) (IEndpoint, error) {
	proto, ok := s.transportProtocols[transport]
	if !ok {
		return nil, ErrUnknownProtocol
	}

	return proto.NewEndpoint(network)
}

// Transport endpoint uses this to find the route, then uses the route to write
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

func (s *Stack) SetRouteTable(table []tcpip.Route) {
	s.routeTable = table
}

// attach so early, the endpoint is even not created
func (s *Stack) CreateNIC(id tcpip.NICID, linkEP tcpip.LinkEndpointID) error {
    return s.createNIC(id, linkEP, true)
}

func (s *Stack) createNIC(id tcpip.NICID, linkEP tcpip.LinkEndpointID, enabled bool) error {
	ep := FindLinkEndpoint(linkEP)
	if ep == nil {
		return tcpip.ErrBadLinkEndpoint
	}

	if _, ok := s.nics[id]; ok {
		return tcpip.ErrDuplicateNICID
	}

	n := newNIC(s, id, ep)

	s.nics[id] = n
	if enabled {
		n.attachLinkEndpoint()
	}

	return nil
}
