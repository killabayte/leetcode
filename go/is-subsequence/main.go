package main

import (
	"fmt"
	"strings"
)

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if strings.Contains(string(t[j]), string(s[i])) {
			i++
		}
		j++
	}
	return i == len(s)
}

func main() {
	s1, t1 := "abc", "ahbgdc"
	fmt.Println(isSubsequence(s1, t1))

	s2, t2 := "axc", "ahbgdc"
	fmt.Println(isSubsequence(s2, t2))
}
