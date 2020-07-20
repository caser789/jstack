package udp

import "testing"

func TestNumber(t *testing.T) {
    p := &protocol{}
    got := p.Number()

    if got != ProtocolNumber {
        t.Fatalf("TestNumber failed. want = %d, got = %d",  ProtocolNumber, got)
    }
}
