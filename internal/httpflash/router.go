package httpflash

import (
	"flash/internal/core"
	"fmt"
	"net/http"
)

type Router struct {
	plugins []core.PluginHTTP
}

func NewRouter(plugins []core.PluginHTTP) *Router {
	return &Router{plugins: plugins}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("chegou no router")
	fmt.Println(r.plugins)
	fmt.Println(req.URL.Path)
	for _, p := range r.plugins {
		if p.Match(req) {
			p.ServeHTTP(w, req)
			return
		}
	}

	http.NotFound(w, req)
}
