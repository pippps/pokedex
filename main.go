package main

import (
	"time"

	"pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		pokeCatched:   make(map[string]pokeapi.PokemonStruct),
	}

	startRepl(cfg)
}
