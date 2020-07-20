package stack

import "errors"
import "github.com/caser789/jstack/tcpip"

const UDPProtocolNumber tcpip.TransportProtocolNumber = 17

var ErrUnknownProtocol = errors.New("unknown protocol")
