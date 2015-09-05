package main

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	// Get memcache client
	memcacheClient := memcache.New("127.0.0.1:11211")

	memcacheClient.Add(&memcache.Item{
		Key:   "title",
		Value: []byte("This is just a test "),
	})

	fetched, _ := memcacheClient.Get("title")

	fmt.Println(string(fetched.Value))

}
