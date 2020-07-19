package stack

import "errors"

const UDPProtocolNumber TransportProtocolNumber = 17

var ErrUnknownProtocol = errors.New("unknown protocol")
