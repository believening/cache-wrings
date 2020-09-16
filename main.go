package main

import (
	"github.com/believening/cache-wrings/cache"
	"github.com/believening/cache-wrings/server"
	_ "github.com/believening/cache-wrings/server/register"
)

func main() {
	c := cache.New("inmemory")
	// serverTyp := "goHttp"
	httpTyp := "httprouter"
	go server.Servers[httpTyp](c).Run()
	server.Servers["tcp"](c).Run()
}
