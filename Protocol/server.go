package Protocol

import (
	"log"
)

type Server interface {
	Run() error
}

func NewServer(param *Param) (ser Server) {
	switch param.Protocol{
	case TCPType:
		ser = NewTCP(param)
	case UDPType:
		ser = NewUDP(param)
	default:
		log.Fatal("Error: Unknown Protocol")
	}

	return ser
}