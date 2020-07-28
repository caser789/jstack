package main

import "log"
import "net"
import "strconv"
import "github.com/caser789/jstack/tcpip"
import "github.com/caser789/jstack/tcpip/network/ipv4"
import "github.com/caser789/jstack/tcpip/stack"
import "github.com/caser789/jstack/tcpip/transport/udp"

func main() {
	log.Println("server started")

	localAddr := "192.168.1.1"

	parsedAddr := net.ParseIP(localAddr)
	log.Printf("parsedAddr is %s", parsedAddr)

	addr := tcpip.Address(parsedAddr.To4())
	proto := ipv4.ProtocolNumber

	localPortName := "9999"
	localPort, _ := strconv.Atoi(localPortName)

	// tunName := "tun0"

	log.Printf("addr is %s", addr)
	log.Printf("proto is %x", proto)
	log.Printf("port is %d", localPort)

	s := stack.New([]string{ipv4.ProtocolName}, []string{udp.ProtocolName})

	log.Printf("stack is %t", s)

	ep, err := s.NewEndpoint(udp.ProtocolNumber, proto)
	if err != nil {
		log.Fatal(err)
	}

	defer ep.Close()

	// if err := ep.Bind(); err != nil {
	//     log.Fatal("Bind failed: ", err)
	// }

	for {
		wq, err := ep.RecvMsg(nil)
		log.Println(wq)
		log.Println(err)
		// ep.SendTo()
	}
}
