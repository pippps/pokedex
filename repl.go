package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
type config struct {
	commands         map[string]cliCommand
	pokeapiClient    pokeapi.Client
	prevLocationsURL *string
	nextLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := cfg.commands[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
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

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapB,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
