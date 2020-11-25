package Protocol

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

const (
	TCP_TYPE = "tcp"
)

type TCPArgs struct {
	Protocol    string
	ListenIP    *string
	ListenPort  *int
	ForwardIP   *string
	ForwardPort *int
}

func NewTCPArgs() *TCPArgs {
	return &TCPArgs{}
}

type TCP struct{}

func NewTCP() Service {
	return &TCP{}
}

func (s *TCP) Run(args interface{}) (err error) {
	tcpArgs := args.(*TCPArgs)
	return tcpArgs.Server()
}

func (s *TCPArgs) Server() (err error) {
	listen, err := net.ListenTCP(
		s.Protocol,
		&net.TCPAddr{
			IP: net.ParseIP(*s.ListenIP), Port: *s.ListenPort,
		})
	if err != nil {
		log.Fatalf("Error to Listen Port: %s", err.Error())
		return
	}

	log.Println("connect init succeed.")
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Fatalf("Error to Accept Traffic: %s", err.Error())
			continue
		}

		log.Printf("connect from: %s", conn.RemoteAddr().String())

		s.Forward(conn)
	}
	return
}

func (s *TCPArgs) Forward(sConn net.Conn) {
	forwardTarget := fmt.Sprintf("%s:%d", *s.ForwardIP, *s.ForwardPort)
	tConn, err := net.Dial(s.Protocol, forwardTarget)
	if err != nil {
		log.Printf("Dial Error: %s", err.Error())
		return
	}

	var wg sync.WaitGroup
	fmt.Println("1")
	go func(sConn net.Conn, tConn net.Conn) {
		wg.Add(1)
		defer wg.Done()
		_, _ = io.Copy(tConn, sConn)
		log.Printf("send: %s -> %s -> %s -> %s", sConn.RemoteAddr(), sConn.LocalAddr(), tConn.RemoteAddr(), tConn.LocalAddr())
		tConn.Close()
	}(sConn, tConn)

	go func(sConn net.Conn, tConn net.Conn) {
		wg.Add(1)
		defer wg.Done()
		_, _ = io.Copy(sConn, tConn)
		log.Printf("accept: %s -> %s -> %s -> %s", tConn.LocalAddr(), tConn.RemoteAddr(), sConn.RemoteAddr(), sConn.LocalAddr())
		sConn.Close()
	}(sConn, tConn)

	wg.Wait()
}
