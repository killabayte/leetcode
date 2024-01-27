package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}

func main() {
	println(strStr("sadbutsad", "sad"))
	println(strStr("leetcode", "leeto"))
	println(strStr("hello", "ll"))
}
