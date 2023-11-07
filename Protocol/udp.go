package protocol

import (
	"fmt"
	"log"
	"net"
)

// UDP UDP转发服务所需参数结构体
type UDP struct {
	Param *Param
}

// NewUDP 创建UDP参数结构体
func NewUDP(param *Param) Server {
	return &UDP{
		Param: param,
	}
}

// Stop 停止UDP转发服务
func (s *UDP) Stop() error {
	return nil
}

// Run 开始UDP转发服务
func (s *UDP) Run() error {
	return s.server()
}

func (s *UDP) server() error {
	conn, err := net.ListenUDP(
		s.Param.Protocol,
		&net.UDPAddr{
			IP:   net.ParseIP(*s.Param.ListenIP),
			Port: *s.Param.ListenPort,
		})
	if err != nil {
		log.Fatalf("Error to Listen Port: %s", err.Error())
		return err
	}

	log.Println("Connect Init Succeed.")

	for {
		data := make([]byte, 1024)
		_, _, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println("failed to read UDP msg because of ", err.Error())
			return err
		}

		s.forward(data)
	}
}

func (s *UDP) forward(data []byte) {
	forwardTarget := fmt.Sprintf("%s:%d", *s.Param.ForwardIP, *s.Param.ForwardPort)
	addr, err := net.ResolveUDPAddr("udp", forwardTarget)
	if err != nil {
		log.Fatalf("Can't resolve address: %s\n", err)
	}

	tConn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Dial Error: %s\n", err)
	}
	defer tConn.Close()

	_, err = tConn.Write(data)
	if err != nil {
		log.Fatalf("Forward Traffic Error: %s", err.Error())
	}
}
