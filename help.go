package main

import "fmt"

func commandHelp(c *config, args ...string) error {
	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
