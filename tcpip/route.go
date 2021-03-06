package tcpip

type Route struct {
	// Destination is the address that must be matched against the masked
	// target address to check if this row is viable.
	Destination Address

	// Mask specifies which bits of the Destination and the target address
	// must match for this row to be viable.
	Mask Address

	// Gateway is the gateway to be used if this row is viable.
	Gateway Address

	// NIC is the id of the nic to be used if this row is viable.
	NIC NICID
}

func (r *Route) Match(addr Address) bool {
	if len(addr) != len(r.Destination) {
		return false
	}

	for i := range r.Destination {
		if (addr[i] & r.Mask[i]) != r.Destination[i] {
			return false
		}
	}

	return true
}
