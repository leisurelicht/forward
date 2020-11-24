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

func (s *TCPArgs) Run() {
	listen, err := net.ListenTCP(
		s.Protocol,
		&net.TCPAddr{
			net.ParseIP(*s.ListenIP),
			*s.ListenPort,
			"",
		})
	if err != nil {
		log.Println("无法监听端口:", err.Error())
		return
	}

	log.Println("已初始化连接, 等待客户端连接")
	s.Server(listen)
}

func (s *TCPArgs) Server(listen *net.TCPListener) {
	for {
		conn, err := listen.AcceptTCP()
		defer conn.Close()
		if err != nil {
			log.Println("接收客户端连接异常:", err.Error())
			continue
		}

		log.Println("客户端连接来自:", conn.RemoteAddr().String())

		s.Forward(conn)
	}
}

func (s *TCPArgs) Forward(sConn net.Conn) {
	forwardTarget := fmt.Sprintf("%s:%d", s.ForwardIP, s.ForwardPort)
	tConn, err := net.Dial(s.Protocol, forwardTarget)
	if err != nil {
		log.Printf("Dial Error: %s", err.Error())
		return
	}

	var wg sync.WaitGroup

	go func(sConn net.Conn, tConn net.Conn) {
		wg.Add(1)
		defer wg.Done()
		io.Copy(tConn, sConn)
		tConn.Close()
	}(sConn, tConn)

	go func(sConn net.Conn, tConn net.Conn) {
		wg.Add(1)
		defer wg.Done()
		io.Copy(sConn, tConn)
		sConn.Close()
	}(sConn, tConn)
	wg.Wait()
}
