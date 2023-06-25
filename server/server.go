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

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	handleGreeting(conn)
}

func handleGreeting(conn net.Conn) {
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
			// Connection between proxy client/server established at this point
			processRequest(buf, n)
		}
	}
}

// func requestHandler(w http.ResponseWriter, _ *http.Request) {
// 	fmt.Printf("request received!\n")
// }

func processRequest(buf []byte, n int) {
	addr := buf[4 : n-2]
	port := binary.BigEndian.Uint16(buf[n-2 : n])

	//  Connect to requested site
	fmt.Printf("Attempting to connect to: %s\n", string(addr)+":"+fmt.Sprint(port))
	serverConn, err := net.Dial("tcp", string(addr)+":"+fmt.Sprint(port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("successfully connected to: %s, on port %d\n", addr, port)
	err = serverConn.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Listen and serve on 8080 as a http server
	//go setupServer()

	// Response to bind to port 8080
	// var port uint16 = 8080
	// portBuf := make([]byte, 2)
	// binary.BigEndian.PutUint16(portBuf, port)

	// bndAddr := net.ParseIP("127.0.0.1")
	// reply := lib.Reply{Version: 0x5, Rep: 0x0, RSV: 0x00, Atyp: 0x01, BndAddr: net.IP.To4(bndAddr), BndPort: portBuf}
	// fmt.Print(reply)
}

// func setupServer() {
// 	http.HandleFunc("/", requestHandler)
// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }
