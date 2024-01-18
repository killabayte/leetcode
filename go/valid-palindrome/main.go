package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	return strings.EqualFold(removeNonLetters(s), reverseString(removeNonLetters(s)))
}

func removeNonLetters(input string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(input, "")
}

func reverseString(input string) string {
	runes := []rune(input)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	input := "A man, a plan, a canal: Panama"
	result := isPalindrome(input)
	fmt.Println(result) // Output: true
}
