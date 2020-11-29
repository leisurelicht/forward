package utils

import (
	"github.com/leisurelicht/forward/Protocol"
	"os"
	"strings"

	kp "gopkg.in/alecthomas/kingpin.v2"
)

const (
	Version = "0.0.1"
)

func ParaseParam() *Protocol.Param {
	param := &Protocol.Param{}

	app := kp.New("forward", "A Simple Traffic Forwarding Tools.")
	app.Author("MuCheng").Version(Version)
	param.ListenIP = app.Flag("listen ip", "the ip is under watch").Short('l').Default("").String()
	param.ListenPort = app.Flag("listen port", "the port is under watch").Short('p').Required().Int()
	param.ForwardIP = app.Flag("forward ip", "forward target ip").Short('F').Required().String()
	param.ForwardPort = app.Flag("forward port", "forward target port").Short('P').Required().Int()

	_ = app.Command("tcp", "TCP Traffic Forward")
	_ = app.Command("udp", "UDP Traffic Forward")

	param.Protocol = strings.ToLower(kp.MustParse(app.Parse(os.Args[1:])))

	return param
}
