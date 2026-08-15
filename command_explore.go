package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *config, args []string) error {
	if len(args) == 0 {
		return errors.New("Please provide an area to explore")
	}

	location := args[0]

	locationDetails, err := c.client.LocationAreaDetails(location)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Printf("You've found the following Pokemon:\n")
	for _, pokemon := range locationDetails.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
