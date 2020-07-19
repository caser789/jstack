package stack

import "testing"

const name = "dummy"
const number = 11

type protocol struct{}

func (*protocol) Number() TransportProtocolNumber {
	return number
}

func TestRegisterTransportProtocol(t *testing.T) {
    p := &protocol{}
    RegisterTransportProtocol(name, p)

    s := New([]string{}, []string{name})

	_, ok := s.(*Stack).transportProtocols[number]

    if want, got := true, ok; want != got {
        t.Fatalf("TestRegisterTransportProtocol failed. want = %t, got = %t",  want, got)
    }
}
