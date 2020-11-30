package Protocol

import (
	"log"
)

type Server interface {
	Run() error
}

func NewServer(param *Param) (ser Server) {
	switch param.Protocol{
	case TCP_TYPE:
		ser = NewTCP(param)
	case UDP_TYPE:
		ser = NewUDP(param)
	default:
		log.Fatal("Error: Unknown Protocol")
	}

	return ser
}