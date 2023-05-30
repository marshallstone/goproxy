package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
		fmt.Println(err)
		return
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte("Received: " + netData))
	}
}
