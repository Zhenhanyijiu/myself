package main

import (
	"fmt"
	"github.com/allegro/bigcache"
	"time"
)

func main() {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))

	cache.Set("my-unique-key", []byte("value111"))

	entry, _ := cache.Get("my-unique-key")
	fmt.Println(string(entry))

}
