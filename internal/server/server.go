package server

import (
	"flash/internal/http"
	"flash/modules/request"
	"fmt"
	"net"
	"os"
)

func Start() {
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Listening on port :4221...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection from", conn.RemoteAddr())

	requestBytes := make([]byte, 1024)
	conn.Read(requestBytes)

	r := request.HttpRequest{}
	req := r.ParseStringToRequest(string(requestBytes))
	response := http.RouteHandler(req)

	conn.Write([]byte(response))
}
