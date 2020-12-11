package main

import (
    "flag"
    "io"
    "log"
    "net"
)

var listenAddr = flag.String("listenAddr", ":666", "listen Addr")

func main() {
	flag.Parse()
	var (
		listen net.Listener
		conn   net.Conn
		err    error
	)
	if listen, err = net.Listen("tcp", *listenAddr); err != nil {
		log.Fatal("Listening Error: ", err)
	}
	defer listen.Close()

	log.Println("Listening on: " + *listenAddr)

	for {
		if conn, err = listen.Accept(); err != nil {
			log.Fatal("Accept Error: ", err)
		}

		go handlerRequest(conn)
	}
}

func handlerRequest(conn net.Conn) {
	defer conn.Close()

	log.Printf("Received message: %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

	if _, err := io.Copy(conn, conn); err != nil {
	    log.Printf("Data Copy Error: %s", err)
    }
}
