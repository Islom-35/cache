package main

import (
	"fmt"
	cache "github.com/homeworkCache/cache"
)

func main() {
	c :=cache.NewCache()
	fmt.Println(c.Set(2, 1))
	fmt.Println(c.Get(2))
	fmt.Println(c.Delete(2))
}
