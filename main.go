package main

import (
	"fmt"
	"time"

	"github.com/DmytroBeliasnyk/in_memory_cache/memory"
)

func main() {
	memCache := memory.GetCache()

	errorsHandler(memCache.Set("login", "password", time.Hour))
	errorsHandler(memCache.Set("login", "password", time.Second))
	errorsHandler(memCache.Set("user1", 1, time.Second))
	errorsHandler(memCache.Set("count", 9, time.Second*5))

	res, err := memCache.Get("login")
	errorsHandler(err)
	fmt.Printf("%v\n\n", res)

	memCache.Delete("login")

	fmt.Println(memCache.String())

	time.Sleep(time.Second * 2)
	fmt.Println(memCache.String())

	time.Sleep(time.Second * 5)
	fmt.Println(memCache.String())
}

func errorsHandler(err error) {
	if err != nil {
		fmt.Printf("%s\n\n", fmt.Errorf("ERROR: %s", err))
	}
}
