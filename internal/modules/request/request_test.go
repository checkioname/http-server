package request

import (
	"reflect"
	"testing"
)

func TestParseRequest(t *testing.T) {
	//Arrange
	rawRequest := "GET /echo/test HTTP/1.1\r\n" +
		"Host: localhost:4221\r\n" +
		"User-Agent: Go-http-client/1.1\r\n" +
		"Accept: */*\r\n\r\n"

	want := HttpRequest{
		Method:        "GET /echo/test HTTP/1.1",
		RequestTarget: "/echo/test",
		HttpVersion:   "HTTP/1.1",
		Host:          "localhost:4221",
		UserAgent:     "Go-http-client/1.1",
		Accept:        "*/*",
	}

	httpRequest := HttpRequest{}

	// Act
	got := httpRequest.ParseStringToRequest(rawRequest)

	// Assert
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ParseRequest() = %v, expected %v", got, want)
	}
}
