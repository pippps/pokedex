package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string
	var inputCleaned []string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input = scanner.Text()
		inputCleaned = cleanInput(input)
		fmt.Printf("Your command was: %s", inputCleaned[0])
		fmt.Print("\n")
	}
}
