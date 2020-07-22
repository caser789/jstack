package stack

import "github.com/caser789/jstack/tcpip"

type NIC struct {
	stack  *Stack
	id     tcpip.NICID
	linkEP ILinkEndpoint
	endpoints   map[NetworkEndpointID]INetworkEndpoint
}

func (n *NIC) findEndpoint(address tcpip.Address) INetworkEndpoint {
	ref := n.endpoints[NetworkEndpointID{address}]
    return ref
}

func newNIC(stack *Stack, id tcpip.NICID, ep ILinkEndpoint) *NIC {
	return &NIC{
		stack:     stack,
		id:        id,
		linkEP:    ep,
		endpoints: make(map[NetworkEndpointID]INetworkEndpoint),
	}
}

// start to deliver packets
func (n *NIC) attachLinkEndpoint() {
	n.linkEP.Attach(n)
}

func (n *NIC) AddAddress(protocol tcpip.NetworkProtocolNumber, addr tcpip.Address) error {
	_, err := n.addAddressLocked(protocol, addr, false)
	return err
}

func (n *NIC) addAddressLocked(protocol tcpip.NetworkProtocolNumber, addr tcpip.Address, replace bool) (*INetworkEndpoint, error) {
	netProto, ok := n.stack.networkProtocols[protocol]
	if !ok {
		return nil, tcpip.ErrUnknownProtocol
	}

	// Create the new network endpoint.
	ep, err := netProto.NewEndpoint(n.id, addr, n, n.linkEP)
	if err != nil {
		return nil, err
	}

	id := *ep.ID()
	n.endpoints[id] = ep

	return ep, nil
}
