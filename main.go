package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "exit",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		inputCleaned := cleanInput(input)
		for i, _ := range inputCleaned {
			cmd, exist := commands[inputCleaned[i]]

			if !exist {
				fmt.Println("Unknown command")
			}
			if err := cmd.callback(); err != nil {
				fmt.Println("Error executing command:", err)
			}
		}

	}
}
