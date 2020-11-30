package main

import (
	"github.com/leisurelicht/forward/Protocol"
	"github.com/leisurelicht/forward/utils"
	"log"
)

func main() {
	param := utils.ParaseParam()
	server := Protocol.NewServer(param)
	if err := server.Run(); err != nil {
		log.Println("Error:", err.Error())
	}
}
