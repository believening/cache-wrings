package main

import (
	"github.com/believening/cache-wrings/cache"
	"github.com/believening/cache-wrings/server"
	_ "github.com/believening/cache-wrings/server/register"
)

func main() {
	c := cache.New("inmemory")
	// serverTyp := "goHttp"
	serverTyp := "httprouter"
	serverFactory := server.Servers[serverTyp]
	serverFactory(c).Run()
}
