package main

import "fmt"

func main() {
	stringSlice := cleanInput("hello world")
	for _, word := range stringSlice {
		fmt.Println(word)
	}
}
