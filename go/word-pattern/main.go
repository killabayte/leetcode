package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	charToWord := make(map[byte]string)
	wordToChar := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		char := pattern[i]
		word := words[i]

		if prevWord, exists := charToWord[char]; exists && prevWord != word {
			return false
		}
		if prevChar, exists := wordToChar[word]; exists && prevChar != char {
			return false
		}

		charToWord[char] = word
		wordToChar[word] = char
	}
	return true
}

func main() {
	pattern := "abba"
	s := "dog cat cat dog"
	result := wordPattern(pattern, s)
	fmt.Println(result)
}
