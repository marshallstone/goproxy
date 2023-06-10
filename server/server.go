package main

import (
	"fmt"
	"log"
	"net"
	"os"

	lib "github.com/marshallstone/goproxy/lib"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("usage: server.go port-number")
		return
	}
	PORT := "127.0.0.1:" + arguments[1]
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	var msg lib.InitPacket
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	if n == 3 {
		fmt.Printf("version: %d\n# auth methods supported: %d\nauth method: %d\n", buf[0], buf[1], buf[2])
	}
}
