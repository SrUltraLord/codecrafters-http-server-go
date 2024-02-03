package main

import (
	"fmt"
	"net"
	"os"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println(" Failed to bind to port 4221")
		os.Exit(1)
	}

	fmt.Println("Server listening on port 4221")

	var conn net.Conn
	conn, err = listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	responseOk := "HTTP/1.1 200 OK\r\n\r\n"

	conn.Write([]byte(responseOk))
}
