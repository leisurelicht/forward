package main

import (
	"flag"
	"log"
	"net"
)

var listenAddr = flag.String("listenAddr", ":777", "listen Addr")

func main() {
	flag.Parse()
	var (
		addr *net.UDPAddr
		conn *net.UDPConn
		err error
	)
	if addr, err = net.ResolveUDPAddr("udp", *listenAddr); err != nil {
		log.Fatal("Can't resolve address: ", err)
	}

	conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("Error listening: ", err)
	}
	defer conn.Close()

	for {
		handleClient(conn)
	}

}

func handleClient(conn *net.UDPConn) {
	data := make([]byte, 1024)
	_, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		log.Printf("Failed to read UDP msg because of %s", err.Error())
		return

	}

	log.Printf("Receivce data From %s", remoteAddr)

	_, err = conn.WriteToUDP(data, remoteAddr)
	if err != nil {
		log.Printf("Write failed, err: %v\n", err)
	}
}