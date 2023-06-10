package main

import (
	"encoding/gob"
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
	msg := lib.RequestMessage{Domain: "https://google.com/", Request: "GET"}

	// Connect to proxy server, create rw
	conn, err := net.Dial("tcp", args[1])
	if err != nil {
		log.Fatal(err)
	}

	enc := gob.NewEncoder(conn)
	// Send a single message over TCP by writing msg to the encoder
	err = enc.Encode(msg)
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()
}
