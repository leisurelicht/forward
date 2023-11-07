package protocol

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

// TCP TCP转发服务所需参数结构体
type TCP struct {
	Param *Param
}

// NewTCP 创建TCP参数结构体
func NewTCP(param *Param) Server {
	return &TCP{
		Param: param,
	}
}

// Stop 停止TCP转发服务
func (s *TCP) Stop() error {
	return nil
}

// Run 开始TCP转发服务
func (s *TCP) Run() (err error) {
	return s.server()
}

func (s *TCP) server() (err error) {
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

	log.Println("Connect Init Succeed.")
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Fatalf("Error to Accept Traffic: %s", err.Error())
			continue
		}

		log.Printf("Connect from: %s", conn.RemoteAddr().String())

		s.forward(conn)
	}
}

func (s *TCP) forward(sConn net.Conn) {
	forwardTarget := fmt.Sprintf("%s:%d", *s.Param.ForwardIP, *s.Param.ForwardPort)
	tConn, err := net.Dial(s.Param.Protocol, forwardTarget)
	if err != nil {
		log.Printf("Dial Error: %s", err.Error())
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func(sConn net.Conn, tConn net.Conn) {
		defer wg.Done()
		_, _ = io.Copy(tConn, sConn)
		log.Printf("send: %s -> %s -> %s -> %s", sConn.RemoteAddr(), sConn.LocalAddr(), tConn.RemoteAddr(), tConn.LocalAddr())
		tConn.Close()
	}(sConn, tConn)

	wg.Add(1)
	go func(sConn net.Conn, tConn net.Conn) {
		defer wg.Done()
		_, _ = io.Copy(sConn, tConn)
		log.Printf("accept: %s -> %s -> %s -> %s", tConn.LocalAddr(), tConn.RemoteAddr(), sConn.RemoteAddr(), sConn.LocalAddr())
		sConn.Close()
	}(sConn, tConn)

	wg.Wait()
}
