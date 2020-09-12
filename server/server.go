package server

import "github.com/believening/cache-wrings/cache"

var Servers map[string]func(c cache.Cache) Server

func init() {
	Servers = make(map[string]func(c cache.Cache) Server)
}

type Server interface {
	Run()
}

func Register(typ string, f func(cache.Cache) Server) {
	Servers[typ] = f
}
