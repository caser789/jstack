package stack

import "github.com/caser789/jstack/tcpip"

type NIC struct {
	stack  *Stack
	id     tcpip.NICID
	linkEP ILinkEndpoint
}

func (n *NIC) findEndpoint(address tcpip.Address) INetworkEndpoint {
    return nil
}

func newNIC(stack *Stack, id tcpip.NICID, ep ILinkEndpoint) *NIC {
	return &NIC{
		stack:     stack,
		id:        id,
		linkEP:    ep,
	}
}

// start to deliver packets
func (n *NIC) attachLinkEndpoint() {
	n.linkEP.Attach(n)
}
