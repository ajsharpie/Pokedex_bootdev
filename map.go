package main

import (
	"fmt"

	pokeapi "github.com/ajsharpie/Pokedex_bootdev/internal/pokeapi"
)

func commandMap(c *config, args ...string) error {
	c.previous = c.next
	c.next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=20&offset=%d", c.offset)
	pokeapi.GetLocationAreas(c.next, c.cache)
	c.offset += 20
	return nil
}

func commandMapb(c *config, args ...string) error {
	if c.offset > 20 {
		c.offset -= 20
	} else {
		c.offset = 0
		fmt.Println("No previous location-areas")
		return nil
	}
	c.next = c.previous
	c.previous = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=20&offset=%d", c.offset)
	pokeapi.GetLocationAreas(c.next, c.cache)
	return nil
}
