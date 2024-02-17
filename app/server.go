package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type Request struct {
	method      string
	path        string
	httpVersion string
	host        string
	userAgent   string
	accept      string
	content     string
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println(" Failed to bind to port 4221")
		os.Exit(1)
	}

	fmt.Println("Server listening on port 4221")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	handleConnection(conn)
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	totalBytesRead, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Could not read request.", err.Error())
	}

	request := parseRequest(totalBytesRead, &buffer)

	var response string

	if request.method == "GET" {
		if request.path == "/" || request.path == "/index.html" {
			response = "HTTP/1.1 200 OK\r\n\r\n"
		} else {
			response = "HTTP/1.1 404 Not Found\r\n\r\n"
		}
	} else {
		response = "HTTP/1.1 404 Not Found\r\n\r\n"
	}

	conn.Write([]byte(response))
	conn.Close()
}

func parseRequest(totalBytesRead int, buffer *[]byte) Request {
	requestBytes := (*buffer)[:totalBytesRead]
	request := strings.Split(string(requestBytes), "\r\n")

	first := strings.Split(request[0], " ")

	return Request{
		method:      first[0],
		path:        first[1],
		httpVersion: first[2],
		host:        strings.Replace(request[1], "Host: ", "", 1),
		userAgent:   strings.Replace(request[2], "User-Agent: ", "", 1),
		accept:      strings.Replace(request[3], "Accept: ", "", 1),
		content:     request[4],
	}
}
