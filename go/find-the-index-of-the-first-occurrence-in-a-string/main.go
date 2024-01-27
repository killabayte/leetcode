package main

import "strings"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return -1
	}

	result := strings.Contains(haystack, needle)
	if result {
		return 0
	}
	return -1
}

func main() {
	println(strStr("sadbutsad", "sad"))
	println(strStr("leetcode", "leeto"))
}
