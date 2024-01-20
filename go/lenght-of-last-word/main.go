package main

import (
	"fmt"
	"strings"
)

func getLastWordLength(s string) int {
	words := strings.Fields(s)
	if len(words) > 0 {
		return len(words[len(words)-1])
	}
	return 0
}

func main() {
	input := "Hello, this is a sample string."

	// Get the last word from the string
	lastWord := getLastWordLength(input)
	fmt.Println("Last Word:", lastWord)
}
