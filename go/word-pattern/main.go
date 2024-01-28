package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(pattern) != len(s) {
		return false
	}

	charToWord := make(map[byte]string)
	wordToChar := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		char := pattern[i]
		word := words[i]

}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
}
