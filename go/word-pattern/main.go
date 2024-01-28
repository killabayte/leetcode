package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	controlMap := make(map[string]string)

	if len(pattern) != len(strings.Split(s, " ")) {
		return false
	}

	for i, v := range strings.Split(s, " ") {
		controlMap[string(pattern[i])] = v
	}
	fmt.Println(controlMap)

	if len(controlMap)%2 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
}
