package main

import (
	"log"

	"github.com/leisurelicht/forward/protocol"
)

func main() {
	param := protocol.ParaseParam()
	server := protocol.NewServer(param)
	if err := server.Run(); err != nil {
		log.Println("Error:", err.Error())
	}
}
