package main

import "fmt"

func commandInspect(cfg *config, pokemonString string) error {
	pokemon, exist := cfg.pokeCatched[pokemonString]
	if !exist {
		fmt.Println("Pokemon not caught")
		return nil
	}
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)

	fmt.Printf("-hp: %d\n-attack: %d\n-defense: %d\n-special-attack: %d\n-special-defense: %d\n-speed: %d\n",
		pokemon.Stats[0].BaseStat, pokemon.Stats[1].BaseStat, pokemon.Stats[2].BaseStat,
		pokemon.Stats[3].BaseStat, pokemon.Stats[4].BaseStat, pokemon.Stats[5].BaseStat)
	fmt.Printf("Types:\n")
	for i := range pokemon.Types {
		fmt.Printf("-%s\n", pokemon.Types[i].Type.Name)
	}
	return nil
}
