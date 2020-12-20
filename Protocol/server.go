package protocol

import (
	"log"
)

// Server 转发服务接口
type Server interface {
	Run() error
	Stop() error
}

// NewServer 根据命令行参数创建不同的转发服务
func NewServer(param *Param) (ser Server) {
	switch param.Protocol {
	case TCPType:
		ser = NewTCP(param)
	case UDPType:
		ser = NewUDP(param)
	default:
		log.Fatal("Error: Unknown Protocol")
	}

	return ser
}
