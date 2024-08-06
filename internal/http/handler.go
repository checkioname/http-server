package http

import (
	"fmt"
	"io/ioutil"
	"crystal/pkg/request"
	"crystal/pkg/response"
	"os"
	"strings"
)

// HandleEcho processes requests to the /echo path
func HandleEcho(req request.HttpRequest) string {
	content, _ := getPathSizeAndContent(req)
	return response.WriteHttpResponse(content)
}

// HandleUserAgent processes requests to the /user-agent path
func HandleUserAgent(req request.HttpRequest) string {
	return response.WriteHttpResponse(req.UserAgent)
}

// HandleFiles processes requests to the /files path
func HandleFiles(req request.HttpRequest) string {
	fileDir := os.Args[2]
	fileName := strings.TrimPrefix(req.RequestTarget, "/files/")

	data, err := ioutil.ReadFile(fileDir + fileName)
	if err != nil {
		return "HTTP/1.1 404 Not Found\r\n\r\n"
	}

	return fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n%s", len(data), data)
}


