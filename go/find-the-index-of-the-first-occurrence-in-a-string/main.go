package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	var neeedleLen int = len(needle)
	var needleFirstChar int

	for _, char := range haystack {
		for _, char2 := range needle {
			if char2 == char {
				neeedleLen--
			}
		}
	}

	if neeedleLen > 0 {
		return -1
	} else {
		return needleFirstChar
	}
}

func main() {
	println(strStr("sadbutsad", "sad"))
	println(strStr("leetcode", "leeto"))
}
