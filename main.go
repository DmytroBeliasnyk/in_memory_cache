package main

import (
	"fmt"

	"github.com/DmytroBeliasnyk/in_memory_cache/memory"
)

func main() {
	memCache := memory.GetCache()

	errorsHandler(memCache.Set("login", "password"))
	errorsHandler(memCache.Set("login2", "password2"))
	errorsHandler(memCache.Set("login2", "password3"))

	res, err := memCache.Get("login")
	errorsHandler(err)
	fmt.Println(res)
	fmt.Println()

	memCache.Delete("login")

	fmt.Println(memCache)
}

func errorsHandler(err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println()
	}
}
