package main

import (
	"log"

	"github.com/believening/cache-wrings/cache"
	"github.com/believening/cache-wrings/server"
)

func main() {
	c := cache.NewInMemCache()
	log.Println("server running")
	server.New(c).Run()
}
