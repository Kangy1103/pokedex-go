package main

import (
	"fmt"
)

func commandHelp(c *config, args []string) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println(" ")
	for name, cmd := range commands {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}
	return nil
}
