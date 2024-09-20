package server

import (
	"flash/internal/http"
	"flash/modules/request"
  "flash/internal/config"
	"fmt"
	"net"
	"os"
)

func Start(config config.Config) {
  port := fmt.Sprintf(":%d", config.Server.Listen)
  listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Failed to bind to port %v", port)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Listening on port :%s...",port)
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
