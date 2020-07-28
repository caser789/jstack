package stack

var networkProtocols = make(map[string]INetworkProtocol)

func RegisterNetworkProtocol(name string, p INetworkProtocol) {
	networkProtocols[name] = p
}
