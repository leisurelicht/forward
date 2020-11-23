package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"sync")

const (
	listenIP    = ""
	listenPort  = 8899
	forwardIP   = "127.0.0.1"
	forwardPort = 8000
	protocol    = "tcp"
)

func Server(listen *net.TCPListener) {
	for {
		conn, err := listen.AcceptTCP()
		defer conn.Close()
		if err != nil {
			log.Println("接收客户端连接异常:", err.Error())
			continue
		}

		log.Println("客户端连接来自:", conn.RemoteAddr().String())

		Forward(conn)
	}
}

func Forward(sConn net.Conn) {
	forwardTarget := fmt.Sprintf("%s:%d", forwardIP, forwardPort)
	tConn, err := net.Dial(protocol, forwardTarget)
	if err != nil {
		log.Printf("Dial Error: %s", err.Error())
		return
	}

	var wg sync.WaitGroup

	go func(sConn net.Conn, tConn net.Conn) {
		wg.Add(1)
		defer wg.Done()
		io.Copy(tConn, sConn)
		tConn.Close()
	}(sConn, tConn)

	go func(sConn net.Conn, tConn net.Conn) {
		wg.Add(1)
		defer wg.Done()
		io.Copy(sConn, tConn)
		sConn.Close()
	}(sConn, tConn)
	wg.Wait()
}

func main() {
	listen, err := net.ListenTCP(protocol, &net.TCPAddr{net.ParseIP(listenIP), listenPort, ""})
	if err != nil {
		log.Println("无法监听端口:", err.Error())
		return
	}

	log.Println("已初始化连接, 等待客户端连接")
	Server(listen)
}
