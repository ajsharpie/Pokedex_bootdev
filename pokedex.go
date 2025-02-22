package main

import (
	"fmt"
)

func commandPokedex(config *config, arg ...string) error {
	fmt.Println("Your Pokedex:")

	if len(config.pokemon) == 0 {
		fmt.Println("  Your Pokedex is empty!")
		return nil
	}
	for _, pokemon := range config.pokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}
