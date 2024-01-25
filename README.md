# Example

package main

import (
	"fmt"

	cache "github.com/islom-35/homework/cache"
)

func main() {
	c := cache.NewCache()
	userID := 1
	name := "Islom"

	text := c.Set(userID, name)
	fmt.Println(text)

	key, val := c.Get(userID)
	fmt.Println(key, val)

	text = c.Delete(userID)
	fmt.Println(text)

}

# Output

saved\n
1 Islom\n
deleted
