package echo

import (
	"flash/internal/core"
	"flash/internal/modules/request"
	"flash/internal/modules/response"
	"net/http"
	"strings"
)

type EchoPlugin struct{}

func (p *EchoPlugin) Name() string {
	return "echo"
}

func (p *EchoPlugin) Match(r *http.Request) bool {
	return r.Method == "GET" && strings.HasPrefix(r.URL.Path, "/echo")
}

func (p *EchoPlugin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rq, _ := request.ParseHttpRequestFromStd(r)
	content := strings.TrimPrefix(rq.RequestTarget, "/echo/")
	httpResponse := response.WriteHttpResponse(content)
	w.Write([]byte(httpResponse))
}

func (p *EchoPlugin) Init(config map[string]interface{}) error {
	return nil
}

func init() {
	core.RegisterModule(&EchoPlugin{})
}
