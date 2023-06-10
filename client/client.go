package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
)

type client_msg struct {
	domain  string
	request string
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("usage: client.go proxy-addr:proxy-port")
	}

	// very simple get request sent to proxy server
	msg := client_msg{domain: "https://google.com/", request: "GET"}

	rw, err := connect(args[1])

	if err != nil {
		log.Fatal(err)
	}
	enc := gob.NewEncoder(rw)
	dec := gob.NewDecoder(rw)
	// Send a single message over TCP by writing msg to the encoder
	enc.Encode(msg)

	// For response
	var resp string
	for {
		dec.Decode(&resp)
		fmt.Printf(resp)
	}
}

func connect(addr string) (*bufio.ReadWriter, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil
}
