package main

import (
	"encoding/gob"
	"fmt"
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
		fmt.Println(err)
		return
	}

	for {
		conn, err := listener.Accept()
		var msg lib.RequestMessage
		dec := gob.NewDecoder(conn)
		err = dec.Decode(&msg)

		if err != nil {
			fmt.Println(err)
			return
		}

		data := fmt.Sprintf("Received msg: {domain: %s, request: %s}", msg.Domain, msg.Request)
		fmt.Printf(data)
	}
}
