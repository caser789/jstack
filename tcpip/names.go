package tcpip

import "errors"

var ErrNotSupported = errors.New("operation not supported")
var ErrNoRoute = errors.New("no route")
var ErrBadLinkEndpoint = errors.New("bad link layer endpoint")
var ErrDuplicateNICID = errors.New("duplicate nic id")
