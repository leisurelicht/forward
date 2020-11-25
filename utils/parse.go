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

func ParaseArgs() *Protocol.TCPArgs {
	tcpArgs := &Protocol.TCPArgs{}

	app := kp.New("forward", "A Simple Traffic Forwarding Tools.")
	app.Author("MuCheng").Version(Version)

	tcp := app.Command("tcp", "TCP Traffic Forward")
	tcpArgs.ListenIP = tcp.Flag("listen ip", "the ip is under watch").Short('l').Default("").String()
	tcpArgs.ListenPort = tcp.Flag("listen port", "the port is under watch").Short('p').Required().Int()
	tcpArgs.ForwardIP = tcp.Flag("forward ip", "forward target ip").Short('F').Required().String()
	tcpArgs.ForwardPort = tcp.Flag("forward port", "forward target port").Short('P').Required().Int()

	tcpArgs.Protocol = strings.ToLower(kp.MustParse(app.Parse(os.Args[1:])))

	return tcpArgs
}
