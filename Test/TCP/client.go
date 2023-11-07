package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

var addr = flag.String("TargetAddr", "127.0.0.1:666", "TargetAddr")

func main() {
	flag.Parse()
	var (
		conn net.Conn
		err  error
	)
	if conn, err = net.Dial("tcp", *addr); err != nil {
		log.Fatalf("Connecting Error: %s \n", err)
	}
	defer conn.Close()

	log.Printf("Connecting to: %s \n", *addr)

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Done()

	go handleWrite(conn)
	go handleRead(conn)

	wg.Wait()
}

func handleWrite(conn net.Conn) {
	buf := make([]byte, 10)
	for {
		// 获取键盘输入。 fmt.Scan --》 结束标记 \n 和 空格
		n, err := os.Stdin.Read(buf) // buf[:n]
		if err != nil {
			fmt.Println("os.Stdin.Read err:", err)
			return
		}
		if string(buf[:3]) == "for" {
			var message string
			for i := 100; i > 0; i-- {
				message += fmt.Sprintf("Hello %d \n", i)
			}
			_ = sendMsg(conn, []byte(message))
			continue
		}
		// 直接将读到键盘输入数据，写到 socket 中，发送给服务器
		_ = sendMsg(conn, buf[:n])
	}
}

func handleRead(conn net.Conn) {
	for {
		_, _ = receiveMsg(conn)
	}
}

func sendMsg(conn net.Conn, msgByte []byte) (err error) {
	// 先发送数据长度
	if _, err = conn.Write(IntToBytes(len(msgByte))); err != nil {
		log.Printf("xxx> [Error] Send Message Size Error: %s \n", err)
		return
	}

	// 发送数据
	if _, err = conn.Write(msgByte); err != nil {
		log.Printf("xxx> [Error] Send Message Content Error: %s \n", err)
		return
	}
	return
}

func receiveMsg(conn net.Conn) (result string, err error) {
	var (
		respByteSize    int
		respByteSizeBuf []byte
		respByteBuf     []byte
	)

	respByteSizeBuf = make([]byte, 8)

	if _, err = conn.Read(respByteSizeBuf); err != nil {
		log.Printf("xxx> [Error] Receive Msg Size Error: %s", err)
		return "", err
	}

	respByteSize = BytesToInt(respByteSizeBuf)

	respByteBuf = make([]byte, respByteSize)
	tmpSize := 0
	for {
		limit, err := conn.Read(respByteBuf[tmpSize:])
		if err != nil {
			log.Printf("xxx> [Error] Read Message Error: %s", err)
			break
		} else if limit == 0 {
			log.Println("---> Read Message Over")
			break
		}
		tmpSize += limit
		if tmpSize == respByteSize {
			log.Println("---> Read All Message")
			break
		}
	}

	log.Printf("---> Response is: %s", string(respByteBuf))

	return string(respByteBuf), nil
}

// 整形转换成字节
func IntToBytes(n int) []byte {
	x := int64(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// 字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int64
	_ = binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
