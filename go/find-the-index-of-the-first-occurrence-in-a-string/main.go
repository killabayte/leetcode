package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	var neeedleLen int = len(needle)

	for i, char := range haystack {
		for j, char2 := range needle {
			if char2 == char {
				if j == neeedleLen-1 {
					return i - j
				}
				continue
			} else {
				break
			}
		}
	}

	if neeedleLen > 0 {
		return -1
	}
}
