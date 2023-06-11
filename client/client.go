package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"

	lib "github.com/marshallstone/goproxy/lib"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("usage: client.go proxy-addr:proxy-port")
	}

	// Create request
	msg := lib.ServerMethod{Version: 0x5, Auth: 1, Method: 0}

	// Connect to proxy server, create rw
	conn, err := net.Dial("tcp", args[1])
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, msg)
	if err != nil {
		log.Fatal(err)
	}
	n, err := conn.Write(buf.Bytes())
	fmt.Printf("%d bytes written to server\n", n)
	if err != nil {
		log.Fatal(err)
	}
	n, err = conn.Read(buf.Bytes())
	fmt.Printf("%d bytes read from server\n", n)

	conn.Close()
}
