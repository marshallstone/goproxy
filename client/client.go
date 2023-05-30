package main

import (
	"bufio"
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
	msg := client_msg{domain: "https://google.com/", request: "GET "}
}

func connect(addr string) (*bufio.ReadWriter, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil
}
