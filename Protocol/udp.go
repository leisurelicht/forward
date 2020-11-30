package Protocol

import (
	"fmt"
	"log"
	"net"
)

type UDP struct {
	Protocol string
	Param    *Param
}

func NewUDP(param *Param) Server {
	return &UDP{
		Protocol: UDP_TYPE,
		Param: param,
	}
}

func (s *UDP) Run() error {
	return s.Server()
}

func (s *UDP) Server() error {
	conn, err := net.ListenUDP(
		s.Protocol,
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

        s.Forward(data)
    }


	return nil
}

func (s *UDP) Forward(data []byte) {
	forwardTarget := fmt.Sprintf("%s:%d", *s.Param.ForwardIP, *s.Param.ForwardPort)
	addr, err := net.ResolveUDPAddr("udp", forwardTarget)
	if err != nil {
		log.Fatalf("Can't resolve address: ", err)
	}

	tConn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Dial Error: ", err)
	}
	defer tConn.Close()

    _, err = tConn.Write(data)
    if err != nil {
        log.Fatalf("Forward Traffic Error: %s", err.Error())
    }

}
