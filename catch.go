package main

import (
	"fmt"
	"math/rand"

	pokeapi "github.com/ajsharpie/Pokedex_bootdev/internal/pokeapi"
)

func commandCatch(c *config, arg ...string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + arg[0]
	pokemon, err := pokeapi.GetPokemon(url, c.cache)
	if err != nil {
		return fmt.Errorf("error fetching Pokemon")
	}
	fmt.Println("Throwing a Pokeball at " + pokemon.Name + "...")
	chance := float64(rand.Intn(pokemon.BaseExperience)) / float64(pokemon.BaseExperience)
	if chance > 0.13 {
		fmt.Println("You caught " + pokemon.Name + "!\n You may now inspect it with the 'inspect' command.")
		c.pokemon[pokemon.Name] = pokemon
	} else {
		fmt.Println("Oh no! " + pokemon.Name + " broke free!")
	}

	return nil
}
