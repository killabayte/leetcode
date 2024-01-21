package main

func canConstruct(ransomNote string, magazine string) bool {
	charCount := make(map[byte]int)

	for i := 0; i < len(magazine); i++ {
		charCount[magazine[i]]++
	}

	// Check if ransomNote can be constructed
	for i := 0; i < len(ransomNote); i++ {
		char := ransomNote[i]
		if charCount[char] <= 0 {
			return false
		}
		charCount[char]--
	}

	return true
}
