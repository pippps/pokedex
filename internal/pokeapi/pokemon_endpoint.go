package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) Pokemon(pokemonName string) (PokemonStruct, error) {
	url := baseURL + "/pokemon/" + pokemonName

	dat, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonStruct{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonStruct{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return PokemonStruct{}, err
		}
		c.cache.Add(url, dat)
	}

	pokemon := PokemonStruct{}
	err := json.Unmarshal(dat, &pokemon)
	if err != nil {
		return PokemonStruct{}, err
	}

	return pokemon, nil
}
