package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args []string) error {
	if len(args) == 0 {
		return errors.New("Which pokemon do you want to catch?")
	}

	pokemon := args[0]

	pokemonDetails, err := c.client.LocationPokemonData(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	result := rand.Intn(pokemonDetails.BaseExperience)
	if result < 40 {
		fmt.Printf("%s was caught!\n", pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}
	return nil
}
