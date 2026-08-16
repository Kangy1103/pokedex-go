package main

import (
	"time"

	"github.com/Kangy1103/pokedex-go/internal/pokeapi"
)

const (
	timeout       time.Duration = 5 * time.Second
	cacheInterval time.Duration = 30 * time.Second
)

func main() {
	c := &config{
		commands: getCommands(),
		client:   pokeapi.NewClient(timeout, cacheInterval),
		pokedex:  make(map[string]pokeapi.PokemonData),
	}
	startRepl(c)
}
