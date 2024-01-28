package main

import (
	"fmt"
)

func wordPattern(pattern string, s string) bool {
	if len(pattern) == 0 || len(s) == 0 {
		return false
	}

	charToWord := make(map[byte]string)
	wordToChar := make(map[string]byte)

}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
}
