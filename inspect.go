package main

import (
	"fmt"
)

func commandInspect(c *config, arg ...string) error {

	if pokemon, ok := c.pokemon[arg[0]]; ok {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Printf("Stats: \n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("   -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("   -%s\n", t.Type.Name)
		}
		return nil
	}
	fmt.Println("You haven't caught that pokemon yet!")
	return fmt.Errorf("pokemon not caught")
}
