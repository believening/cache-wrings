package httprouter

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/believening/cache-wrings/cache"
	"github.com/believening/cache-wrings/server"
)

var (
	port = ":12345"
	typ  = "httprouter"
)

type httpRouterServer struct {
	cache.Cache
}

func init() {
	server.Register(typ, New)
}

func New(c cache.Cache) server.Server {
	return &httpRouterServer{c}
}

func (s *httpRouterServer) Run() {
	router := httprouter.New()
	c := &cacheHandler{s}
	router.Handle(http.MethodGet, "/cache/:key", c.GetHandle)
	router.Handle(http.MethodPut, "/cache/:key", c.PutHandle)
	router.Handle(http.MethodDelete, "/cache/:key", c.DeleteHandle)
	router.Handler(http.MethodGet, "/status", &statusHandler{s})
	_ = http.ListenAndServe(port, router)
}
