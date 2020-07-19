package stack

import "testing"
// import "log"

const name = "dummy"
const number = 11

type protocol struct{
    endpoint IEndpoint
}
type endpoint struct{}

func (*protocol) Number() TransportProtocolNumber {
	return number
}

func (p *protocol) NewEndpoint(netProto NetworkProtocolNumber) (IEndpoint, error) {
	return p.endpoint, nil
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

func TestNewEndpoint(t *testing.T) {
    want := &endpoint{}
    p := &protocol{endpoint: want}
    RegisterTransportProtocol(name, p)

    s := New([]string{}, []string{name})

    got, _ := s.NewEndpoint(number, 2)

    if want != got {
        t.Fatalf("TestNewEndpoint failed. want = %t, got = %t",  want, got)
    }
}
