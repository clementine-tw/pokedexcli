package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(location string) (Location, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		cached := Location{}
		err := json.Unmarshal(val, &cached)
		if err != nil {
			return cached, err
		}
		return cached, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}
	c.cache.Add(url, data)

	pokemons := Location{}
	err = json.Unmarshal(data, &pokemons)
	if err != nil {
		return pokemons, err
	}

	return pokemons, nil
}
