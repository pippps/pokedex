package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\n" +
		"Usage:\n" +
		"\n" +
		"help: Displays a help message\n" +
		"exit: Exit the Pokedex")
	return nil
}
