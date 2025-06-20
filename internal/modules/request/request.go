package request

import (
	"net/http"
)

type HttpRequest struct {
	Method        string
	RequestTarget string
	HttpVersion   string
	Host          string
	UserAgent     string
	Accept        string
	Body          string
}

func ParseHttpRequestFromStd(r *http.Request) (*HttpRequest, error) {
	return &HttpRequest{
		Method:        r.Method + " " + r.RequestURI + " " + r.Proto,
		RequestTarget: r.URL.Path,
		HttpVersion:   r.Proto,
		Host:          r.Host,
		UserAgent:     r.UserAgent(),
		Accept:        r.Header.Get("Accept"),
		Body:          "",
	}, nil
}
