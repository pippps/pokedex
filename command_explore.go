package main

import "fmt"

func commandExplore(cfg *config, locationArea string) error {
	pokemonsResp, err := cfg.pokeapiClient.ListPokemon(locationArea)
	if err != nil {
		return err
	}
	fmt.Println("Exploring pastoria-city-area...")
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonsResp.PokemonList {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
