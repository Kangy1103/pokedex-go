package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, args []string) error {
	if len(args) == 0 {
		return errors.New("please provide a Pokemon name")
	}

	pokemon := args[0]

	if poke, ok := c.pokedex[pokemon]; ok {
		fmt.Printf("Name: %s\n", poke.Name)
		fmt.Printf("Height: %d\n", poke.Height)
		fmt.Printf("Weight: %d\n", poke.Weight)
		fmt.Printf("Stats:")
		for _, stat := range poke.Stats {
			fmt.Printf("	-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:")
		for _, pokeType := range poke.Types {
			fmt.Printf("	-%s", pokeType.Type.Name)
		}
	}

	return nil
}
