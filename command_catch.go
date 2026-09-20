package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, pokemonName string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemon, err := cfg.pokeapiClient.Pokemon(pokemonName)
	if err != nil {
		return err
	}
	randomNum := rand.Intn(600)
	if randomNum > pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokeCatched[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}
