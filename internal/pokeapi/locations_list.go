package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locations RespShallowLocations
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return RespShallowLocations{}, err
	}
	cached, err := json.Marshal(locations)
	if err != nil {
		return RespShallowLocations{}, err
	}
	c.cache.Add(url, cached)

	return locations, nil
}
