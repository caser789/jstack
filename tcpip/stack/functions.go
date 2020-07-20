package stack

import "github.com/caser789/jstack/tcpip"

func New(network []string, transport []string) IStack {
    s := &Stack{
		transportProtocols: make(map[tcpip.TransportProtocolNumber]ITransportProtocol),
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

