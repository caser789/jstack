package stack

import "testing"
import "log"

func TestStack(t *testing.T) {
    s := New()
    ep := s.NewEndpoint()

    log.Println(ep)
}
