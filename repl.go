package main

import "strings"

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
