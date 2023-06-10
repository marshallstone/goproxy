package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

type client_msg struct {
	domain  string
	request string
}

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
		var msg client_msg
		rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
		dec := gob.NewDecoder(rw)
		err := dec.Decode(&msg)

		if err != nil {
			fmt.Println(err)
			return
		}

		data := fmt.Sprintf("Received msg: {domain: %s, request: %s}", msg.domain, msg.request)
		conn.Write([]byte(data))
	}
}
