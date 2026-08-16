package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// LocationPokemonData Mostly copied from location_list.go
func (c *Client) LocationPokemonData(selectedPokemon string) (PokemonData, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + selectedPokemon
	pokemon := PokemonData{}

	if cacheData, found := c.cache.Get(url); found {
		if err := json.Unmarshal(cacheData, &pokemon); err != nil {
			return pokemon, err
		}
		return pokemon, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemon, err
	}
	locale, err := c.httpClient.Do(request)
	if err != nil {
		fmt.Println("Unable to get Pokemon details...")
		return pokemon, err
	}
	data, err := io.ReadAll(locale.Body)
	if err != nil {
		return pokemon, err
	}
	c.cache.Add(url, data)

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return pokemon, err
	}
	return pokemon, nil
}
