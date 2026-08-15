package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationAreasList(pageURL *string) (AreaLocations, error) {
	url := "https://pokeapi.co/api/v2/location-area/"
	areas := AreaLocations{}
	if pageURL == nil {
		pageURL = &url
	}

	if cacheData, found := c.cache.Get(*pageURL); found {
		if err := json.Unmarshal(cacheData, &areas); err != nil {
			return areas, err
		}
		return areas, nil
	}

	request, err := http.NewRequest("GET", *pageURL, nil)
	if err != nil {
		return areas, err
	}
	locations, err := c.httpClient.Do(request)
	if err != nil {
		fmt.Println("Unable to get location areas...")
		return areas, err
	}
	data, err := io.ReadAll(locations.Body)
	if err != nil {
		return areas, err
	}
	c.cache.Add(*pageURL, data)

	if err := json.Unmarshal(data, &areas); err != nil {
		return areas, err
	}
	return areas, nil
}
