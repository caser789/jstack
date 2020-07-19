package stack


const UDPProtocolNumber TransportProtocolNumber = 17

var transportProtocols = make(map[string]ITransportProtocol)
