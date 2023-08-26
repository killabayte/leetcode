package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	stack := []rune{}
	bracketPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if strings.Contains("([{", string(char)) {
			stack = append(stack, char)
		} else if strings.Contains(")]}", string(char)) {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != bracketPairs[char] {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
}
