package main

import (
	pokeapi "github.com/ajsharpie/Pokedex_bootdev/internal/pokeapi"
)

func commandExplore(c *config, arg ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + arg[0]
	pokeapi.GetPokemonFromLoc(url, c.cache)
	return nil
}
