package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
)

var addr = flag.String("TargetAddr", "127.0.0.1:666", "TargetAddr")

func main() {
	flag.Parse()
	var (
		conn net.Conn
		err  error
	)
	if conn, err = net.Dial("tcp", *addr); err != nil {
		log.Fatalf("Connecting Error: %s \n", err)
	}
	defer conn.Close()

	log.Printf("Connecting to: %s \n", *addr)

	var wg sync.WaitGroup
	wg.Add(2)

	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)

	wg.Wait()
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 3; i > 0; i-- {
		message := fmt.Sprintf("Hello %d \n", i)
		if _, err := conn.Write([]byte(message)); err != nil {
			log.Printf("Send Message Error: %s \n", err)
			break
		}
	}
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		log.Fatalf("Read Message Error: %s \n", err)
	}
	log.Println(string(buf[:reqLen-1]))
}
