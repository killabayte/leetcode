package main

import "fmt"

func productExceptSelf(nums []int) []int {
	length := len(nums)

	// Initialize prefix and suffix arrays
	prefix := make([]int, length)
	suffix := make([]int, length)

	// Calculate prefix products
	prefixProduct := 1
	for i := 0; i < length; i++ {
		prefix[i] = prefixProduct
		prefixProduct *= nums[i]
	}

	// Calculate suffix products
	suffixProduct := 1
	for i := length - 1; i >= 0; i-- {
		suffix[i] = suffixProduct
		suffixProduct *= nums[i]
	}

	// Calculate the result array
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = prefix[i] * suffix[i]
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println(result)
}
