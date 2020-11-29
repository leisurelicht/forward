package Protocol

import (
    "fmt"
    "log"
    "net"
    "os"
)

type UDP struct {
	Protocol string
	Param    *Param
}

func NewUDP() Service {
	return &UDP{
		Protocol: UDP_TYPE,
	}
}

func (s *UDP) Run(args interface{}) error {
	s.Param = args.(*Param)
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

	log.Println("connect init succeed.")


	for {
        data := make([]byte, 1024)
        n, remoteAddr, err := conn.ReadFromUDP(data)
        if err != nil {
            fmt.Println("failed to read UDP msg because of ", err.Error())
            return err
        }
        fmt.Println(n, remoteAddr)

        s.Forward(data)
    }


	return nil
}

func (s *UDP) Forward(data []byte) {
	forwardTarget := fmt.Sprintf("%s:%d", *s.Param.ForwardIP, *s.Param.ForwardPort)
	addr, err := net.ResolveUDPAddr("udp", forwardTarget)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}

	tConn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	defer tConn.Close()

    _, err = tConn.Write(data)
    if err != nil {
        fmt.Println("failed:", err)
        os.Exit(1)
    }

}
