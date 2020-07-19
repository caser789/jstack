package stack

type Stack struct {
    transportProtocols map[TransportProtocolNumber]ITransportProtocol
}

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

func (s *Stack) NewEndpoint(transport TransportProtocolNumber, network NetworkProtocolNumber) (IEndpoint, error) {
    proto := s.transportProtocols[transport]

    return proto.NewEndpoint(network)
}
