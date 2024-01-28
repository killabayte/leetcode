package main

import "strings"

func wordPattern(pattern string, s string) bool {
	controlMap := make(map[string]string)

	if len(pattern) != len(strings.Split(s, " ")) {
		return false
	}

}
