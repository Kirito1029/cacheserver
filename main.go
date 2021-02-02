package main

import (
	"cache-server/pkg/cache"
	"cache-server/pkg/hashtable"
)

func main() {
	hashtable.New()
	cache.Start()
}
