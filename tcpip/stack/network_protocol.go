package stack

var networkProtocols   = make(map[string]NetworkProtocol)

func RegisterNetworkProtocol(name string, p NetworkProtocol) {
	networkProtocols[name] = p
}
