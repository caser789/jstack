package stack

// global
var transportProtocols = make(map[string]ITransportProtocol)

func RegisterTransportProtocol(name string, p ITransportProtocol) {
	transportProtocols[name] = p
}
