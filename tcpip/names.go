package tcpip

import "errors"

var ErrNotSupported = errors.New("operation not supported")
var ErrNoRoute = errors.New("no route")
var ErrBadLinkEndpoint = errors.New("bad link layer endpoint")
var ErrDuplicateNICID = errors.New("duplicate nic id")
var ErrUnknownProtocol= errors.New("unknown protocol")
var ErrUnknownNICID = errors.New("unknown nic id")
