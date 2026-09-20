package main

import "fmt"

func commandPokedex(cfg *config, parameter string) error {
	fmt.Println("Your pokedex:")
	if len(cfg.pokeCatched) == 0 {
		fmt.Println("Catch some pokemon first!\nYour pokedex is empty.")
	}
	for k := range cfg.pokeCatched {
		fmt.Printf("-%s\n", k)
	}
	return nil
}
