package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/ajsharpie/Pokedex_bootdev/internal/pokeapi"
)

func StartRepl() {
	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)

	config := &config{
		next:     "https://pokeapi.co/api/v2/location-area/?limit=20",
		previous: "",
		offset:   0,
		cache:    pokeapi.NewCache(60),
		pokemon:  map[string]pokeapi.Pokemon{},
	}

	for i := 0; i >= 0; i++ {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]
		if command == "help" {
			commandHelp(config)
		} else {
			if command, ok := commands[command]; ok {
				if len(cleaned) > 1 {
					command.callback(config, cleaned[1])
				} else {
					command.callback(config)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the pokedex",
		callback:    commandExit,
	},
	"map": {
		name:        "map",
		description: "List the next 20 location-areas",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "List the previous 20 location-areas",
		callback:    commandMapb,
	},
	"printcache": {
		name:        "printcache",
		description: "Print the current cache",
		callback:    commandPrintCache,
	},
	"explore": {
		name:        "explore <arg>",
		description: "explore the named location area",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch <arg>",
		description: "catch the named pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect <arg>",
		description: "inspect the named pokemon",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "List all the pokemon in the pokedex",
		callback:    commandPokedex,
	},
}

type config struct {
	next     string
	previous string
	offset   int
	cache    *pokeapi.Cache
	pokemon  map[string]pokeapi.Pokemon
}

func commandExit(c *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.Trim(text, " ")
	words := strings.Split(text, " ")
	return words
}
