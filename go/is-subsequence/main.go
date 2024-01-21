package main

import "strings"

func isSubsequence(s string, t string) bool {
	return strings.Contains(t, s)
}
