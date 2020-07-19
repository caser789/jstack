package main

import "log"

const IPv4ProtocolName = "ipv4"
const UDPProtocolName = "udp"


type UDPEndpoint struct {}

func main() {
    log.Println("server started")

	s := New()

	ep, err := s.NewEndpoint()
	if err != nil {
		log.Fatal(err)
	}

	// defer ep.Close()

	// if err := ep.Bind(); err != nil {
	// 	log.Fatal("Bind failed: ", err)
	// }

    // for {

	// 	n, wq, err := ep.RecvFrom()
	// 	ep.SendTo()
    // }
}
