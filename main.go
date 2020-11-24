package main

import (
	"github.com/leisurelicht/forward/Protocol"
	"github.com/leisurelicht/forward/utils"
	"log"
)

var ser Protocol.Service

func main() {
	args := utils.ParaseArgs()

	switch args.Protocol {
	case Protocol.TCP_TYPE:
		ser = Protocol.NewTCP()
	//case Protocol.UDP_TYPE:
	//	ser = Protocol.NewUDP()
	}

	if err := ser.Run(args); err != nil {
		log.Println("Error:", err.Error())
	}
}
