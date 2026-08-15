package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kangy1103/pokedex-go/internal/pokeapi"
)

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	return strings.Fields(lowercaseText)
}

type config struct {
	commands map[string]cliCommand
	client   pokeapi.Client
	next     *string
	previous *string
	explore  *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, args []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show this help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show the first 20 of Pokemon locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 Pokemon locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a specific location",
			callback:    commandExplore,
		},
	}
}

func startRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanText := cleanInput(text)
		if len(cleanText) == 0 {
			continue
		}
		command := cleanText[0]
		args := cleanText[1:]

		cmd, ok := c.commands[command]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := cmd.callback(c, args)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
