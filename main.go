package main

import (
	"github.com/leisurelicht/forward/protocol"
	"github.com/leisurelicht/forward/utils"
	"log"
)

func main() {
	param := utils.ParaseParam()
	server := protocol.NewServer(param)
	if err := server.Run(); err != nil {
		log.Println("Error:", err.Error())
	}
}
