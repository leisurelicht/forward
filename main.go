package main

import (
	"github.com/leisurelicht/forward/Protocol"
	"github.com/leisurelicht/forward/utils"
)


func main() {
	args := utils.ParaseArgs()
	switch args.Protocol {
	case Protocol.TCP_TYPE:
		args.Run()
	}
}
