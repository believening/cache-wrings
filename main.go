package main

import (
	"cache-wings/cache"
	"cache-wings/server"
)

func main() {
	c := cache.NewInMemCache()
	server.New(c).Run()
}
