package Protocol

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type TCP struct {
	Protocol string
	Param    *Param
}

func NewTCP() Service {
	return &TCP{
		Protocol: TCP_TYPE,
	}
}

func (s *TCP) Run(args interface{}) (err error) {
	s.Param = args.(*Param)
	return s.Server()
}

func (s *TCP) Server() (err error) {
	listen, err := net.ListenTCP(
		s.Param.Protocol,
		&net.TCPAddr{
			IP:   net.ParseIP(*s.Param.ListenIP),
			Port: *s.Param.ListenPort,
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

func (s *TCP) Forward(sConn net.Conn) {
	forwardTarget := fmt.Sprintf("%s:%d", *s.Param.ForwardIP, *s.Param.ForwardPort)
	tConn, err := net.Dial(s.Param.Protocol, forwardTarget)
	if err != nil {
		log.Printf("Dial Error: %s", err.Error())
		return
	}

	var wg sync.WaitGroup
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
