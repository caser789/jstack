package stack

func New(network []string, transport []string) IStack {
    s := &Stack{
		transportProtocols: make(map[TransportProtocolNumber]ITransportProtocol),
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

