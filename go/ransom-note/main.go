package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	charCount := make(map[byte]int)

	for i := 0; i < len(magazine); i++ {
		charCount[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		char := ransomNote[i]
		if charCount[char] <= 0 {
			return false
		}
		charCount[char]--
	}

	return true
}

func main() {
	ransomNote1, magazine1 := "a", "b"
	fmt.Println(canConstruct(ransomNote1, magazine1)) // Output: false

	// Example 2
	ransomNote2, magazine2 := "aa", "ab"
	fmt.Println(canConstruct(ransomNote2, magazine2)) // Output: false

	// Example 3
	ransomNote3, magazine3 := "aa", "aab"
	fmt.Println(canConstruct(ransomNote3, magazine3)) // Output: true
}
