/*
Package pokeapi: The pokeapi is used to perform all the background actions of the pokedex-go cli application
There ain't really a top level file so to shut my nvim lsp up i've put this here
I don't know what's doing it but it's pretty annoying and flagging any comments without "package x" or "<func name> some text here..."
*/
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// LocationAreaDetails Mostly copied from location_list.go
func (c *Client) LocationAreaDetails(location string) (AreaDetails, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + location
	area := AreaDetails{}

	if cacheData, found := c.cache.Get(url); found {
		if err := json.Unmarshal(cacheData, &area); err != nil {
			return area, err
		}
		return area, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return area, err
	}
	locale, err := c.httpClient.Do(request)
	if err != nil {
		fmt.Println("Unable to get location details...")
		return area, err
	}
	data, err := io.ReadAll(locale.Body)
	if err != nil {
		return area, err
	}
	c.cache.Add(url, data)

	if err := json.Unmarshal(data, &area); err != nil {
		return area, err
	}
	return area, nil
}
