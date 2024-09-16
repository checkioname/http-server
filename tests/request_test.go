package request

import (
	"reflect"
	"testing"
  "crystal/pkg/request"
)

func TestParseRequest(t *testing.T) {
	rawRequest := "GET /echo/test HTTP/1.1\r\n" +
		            "Host: localhost:4221\r\n" +
            		"User-Agent: Go-http-client/1.1\r\n" +
            		"Accept: */*\r\n\r\n"

	expected := request.HttpRequest{
		Method:        "GET /echo/test HTTP/1.1",
		RequestTarget: "/echo/test",
		HttpVersion:   "HTTP/1.1",
		Host:          "localhost:4221",
		UserAgent:     "Go-http-client/1.1",
		Accept:        "*/*",
	}

	// Chama a função ParseRequest com a requisição bruta
	result := request.ParseRequest(rawRequest)
  
	// Verifica se o resultado é igual ao esperado
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseRequest() = %v, expected %v", result, expected)
	}
}

