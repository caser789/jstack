package stack

type Stack struct {
    transportProtocols map[TransportProtocolNumber]ITransportProtocol
}

func (s *Stack) NewEndpoint(transport TransportProtocolNumber, network NetworkProtocolNumber) (IEndpoint, error) {
	proto, ok := s.transportProtocols[transport]
	if !ok {
		return nil, ErrUnknownProtocol
	}

	return proto.NewEndpoint(network)
}
