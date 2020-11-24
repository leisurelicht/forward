package utils

import (
    "os"
    "strings"
    "github.com/leisurelicht/forward/Protocol"

    kp "gopkg.in/alecthomas/kingpin.v2"
)

const (
	Version  = "0.0.1"
)

func ParaseArgs() *Protocol.TCPArgs {
	args := &Protocol.TCPArgs{}

	app := kp.New("forward", "A Simple Traffic Forwarding Tools.")
	app.Author("MuCheng").Version(Version)

	tcp := app.Command("tcp", "TCP Traffic Forward")
	args.ListenIP = tcp.Flag("listen ip", "the ip is under watch").Short('l').Default("").String()
	args.ListenPort = tcp.Flag("listen port", "the port is under watch").Short('p').Required().Int()
	args.ForwardIP = tcp.Flag("forward ip", "forward target ip").Short('F').Required().String()
	args.ForwardPort = tcp.Flag("forward port", "forward target port").Short('P').Required().Int()

	args.Protocol = strings.ToLower(kp.MustParse(app.Parse(os.Args[1:])))

	return args
}
