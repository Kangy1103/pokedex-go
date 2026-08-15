package main

import (
	"errors"
	"fmt"
)

func commandMap(c *config, args []string) error {
	areas, err := c.client.LocationAreasList(c.next)
	if err != nil {
		return err
	}

	c.next = areas.Next
	c.previous = areas.Previous
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(c *config, args []string) error {
	if c.previous == nil {
		return errors.New("You can't do that Dave...")
	}
	areas, err := c.client.LocationAreasList(c.previous)
	if err != nil {
		return err
	}

	c.next = areas.Next
	c.previous = areas.Previous
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}
