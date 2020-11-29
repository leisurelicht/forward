package main

import (
	"github.com/leisurelicht/forward/Protocol"
	"github.com/leisurelicht/forward/utils"
	"log"
)

var ser Protocol.Service

func main() {
	param := utils.ParaseParam()

	switch param.Protocol {
	case Protocol.TCP_TYPE:
		ser = Protocol.NewTCP()
	case Protocol.UDP_TYPE:
		ser = Protocol.NewUDP()
	}

	if err := ser.Run(param); err != nil {
		log.Println("Error:", err.Error())
	}
}
