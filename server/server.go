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
	// Use the same buffer for read/write operations?
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)

		if err != nil {
			log.Fatal(err)
		}
		// Initial client message, start handhsake
		if n == 3 {
			initMsg := lib.ServerMethod{Version: buf[0], Auth: buf[1], Method: buf[2]}

			if initMsg.Version != 5 {
				log.Fatal("socks version < 5 not supported.\n")
			}

			// Send method with no auth (no subnegotiation required)
			resp := lib.ClientMethod{Version: 0x5, Method: 0x0}
			respBuf := new(bytes.Buffer)
			err = binary.Write(respBuf, binary.BigEndian, resp)
			if err != nil {
				log.Fatal(err)
			}

			conn.Write(respBuf.Bytes())
			if err != nil {
				log.Fatal(err)
			}
		} else {
			// Handle requests next
			// Form TCP connection to forward requests
			fmt.Printf("%d bytes read\n", n)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
