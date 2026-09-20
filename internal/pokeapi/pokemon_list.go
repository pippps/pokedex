package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(location_area string) (respPokemon, error) {
	url := baseURL + "/location-area/" + location_area

	dat, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return respPokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return respPokemon{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return respPokemon{}, err
		}
		c.cache.Add(url, dat)
	}

	pokemonsResp := respPokemon{}
	err := json.Unmarshal(dat, &pokemonsResp)
	if err != nil {
		return respPokemon{}, err
	}

	return pokemonsResp, nil
}
