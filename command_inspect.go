package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, args []string) error {
	if len(args) == 0 {
		return errors.New("Please provide a Pokemon name")
	}

	pokemon := args[0]

	if poke, ok := c.client.LocationAreaDetails(pokemon); ok {
	}

	return nil
}
