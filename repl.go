package main

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	var tempWord string
	var wordSlice []string
	text = strings.ToLower(text)
	for _, ch := range text {
		if ch == ' ' && tempWord != "" {
			wordSlice = append(wordSlice, tempWord)
			tempWord = ""
		} else if ch != ' ' {
			if tempWord == "" {
				tempWord = string(ch)
			} else {
				tempWord += string(ch)
			}
		}

	}
	if tempWord != "" {
		wordSlice = append(wordSlice, tempWord)
	}
	return wordSlice
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\n" +
		"Usage:\n" +
		"\n" +
		"help: Displays a help message\n" +
		"exit: Exit the Pokedex")
	return nil
}
