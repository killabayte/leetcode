package main

import (
	"fmt"
	"strings"
)

func isSubsequence(s string, t string) bool {
	return strings.Contains(t, s)
}
func main() {
	result := isSubsequence("abc", "ahbgdc")
	fmt.Println(result)
}
