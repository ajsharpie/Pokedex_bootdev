package main

import (
	"fmt"
)

func commandPrintCache(c *config, args ...string) error {
	fmt.Println("Current Cache:")
	for key, value := range c.cache.Entries {
		fmt.Printf("%v: %v\n", key, value)
	}
	return nil
}
