package main

import (
	"errors"
	"fmt"
)

func commandPokedex(c *config, args []string) error {
	if len(c.pokedex) == 0 {
		return errors.New("your Pokedex is empty")
	}
	fmt.Println("Your Pokedex:")
	for pokemon := range c.pokedex {
		fmt.Printf("  -%s\n", pokemon)
	}
	return nil
}
