package gohttp

import (
	"net/http"

	"github.com/believening/cache-wrings/cache"
	"github.com/believening/cache-wrings/server"
)

var (
	port = ":12345"
	typ  = "goHttp"
)

type goHTTPServer struct {
	cache.Cache
}

func init() {
	server.Register(typ, New)
}

func New(c cache.Cache) server.Server {
	return &goHTTPServer{c}
}

func (s *goHTTPServer) Run() {
	s.Listen()
}

func (s *goHTTPServer) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHanler())
	_ = http.ListenAndServe(port, nil)
}

func (s *goHTTPServer) cacheHandler() http.Handler {
	return &cacheHanler{s}
}

func (s *goHTTPServer) statusHanler() http.Handler {
	return &statusHandler{s}
}
