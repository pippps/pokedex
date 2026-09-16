package main

func cleanInput(text string) []string {
	var tempWord string
	var wordSlice []string
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
