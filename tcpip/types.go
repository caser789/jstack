package tcpip

type TransportProtocolNumber uint32
type NetworkProtocolNumber uint32

type Address string

type FullAddress struct {
	Addr Address
	Port uint16
    NIC NICID
}

type NICID int32
