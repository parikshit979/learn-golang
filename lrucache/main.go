package main

import (
	"fmt"
	"log"

	"github.com/learn-golang/lrucache/cache"
)

func main() {
	fmt.Println("Starting LRU Cache")

	cache := cache.NewCache(3)
	cache.Set("1", "one")
	cache.Set("2", "two")
	cache.Set("3", "three")
	log.Printf("Get for key(1): %s", cache.Get("1"))
}
