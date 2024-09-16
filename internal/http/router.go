package http

import (
	"crystal/pkg/request"
	"crystal/pkg/response"
	"net/http"
	"strings"
)


type Route struct {
  Method string // GET, POST, PUT, DELETE
  Path string  // caminho request
  Handler []HandlerFunc  //Middleware
  Group string // grupo para o qual a rota pertence
  IsTerminal bool // se verdadeiro nenhuma rota processa depois

}

type Router struct {
  Routes []Route
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func RouteHandler(req request.HttpRequest) string {
	if !isValidPath(req.Method) {
		return "HTTP/1.1 404 Not Found\r\n\r\n"
	}

	if strings.Contains(req.Method, "GET /echo") {
		content, _ := getPathSizeAndContent(req)
		return response.WriteHttpResponse(content)
	}

	if req.RequestTarget == "/user-agent" {
		return response.WriteHttpResponse(req.UserAgent)
	}

	if strings.Contains(req.Method, "GET /files/") {
		// Lógica para lidar com arquivos
	}

	return "HTTP/1.1 200 OK\r\n\r\n"
}

func isValidPath(method string) bool{
    return method == "GET / HTTP/1.1" || strings.Contains(method,"GET /echo/") || strings.Contains(method,"GET /user-agent") || strings.Contains(method,"GET /files")
}

func getPathSizeAndContent(req request.HttpRequest) (string,int){
  content := strings.Split(req.RequestTarget,"/")
  lastElement := content[len(content)-1]
  var contentLength = len(lastElement)

  return lastElement, contentLength
}
