package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocation(pageURL *string) (respShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	dat, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return respShallowLocations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return respShallowLocations{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return respShallowLocations{}, err
		}
		c.cache.Add(url, dat)
	}

	locationsResp := respShallowLocations{}
	err := json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return respShallowLocations{}, err
	}

	return locationsResp, nil
}
