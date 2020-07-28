package tcpip

import "testing"
import "log"

func TestMatch(t *testing.T) {
	r := &Route{
		Destination: Address([]byte{192, 168, 0, 0}),
		Mask:        Address([]byte{255, 255, 0, 0}),
		Gateway:     "127.0.0.1",
		NIC:         1,
	}

	res := r.Match(Address([]byte{1, 1, 1, 1}))
	log.Println(res)

	res = r.Match(Address([]byte{192, 168, 1, 1}))
	log.Println(res)

	res = r.Match(Address([]byte{192, 168, 0, 1}))
	log.Println(res)
}
