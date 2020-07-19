package stack

type TransportProtocolNumber uint32
type NetworkProtocolNumber uint32

type IEndpoint interface {}

type IStack interface {
	// NewEndpoint(transport TransportProtocolNumber, network NetworkProtocolNumber) (IEndpoint, error)
}

type ITransportProtocol interface {
	// NewEndpoint(netProto NetworkProtocolNumber) (IEndpoint, error)
    Number() TransportProtocolNumber
}
