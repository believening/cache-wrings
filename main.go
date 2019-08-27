package main

import (
	"gtihub.com/believening/cache-wrings/cache"
	"gtihub.com/believening/cache-wrings/server"
)

func main() {
	c := cache.NewInMemCache()
	server.New(c).Run()
}
