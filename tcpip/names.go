package tcpip

import "errors"

var ErrNotSupported = errors.New("operation not supported")
var ErrNoRoute = errors.New("no route")
