package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	uniqueIndex := 2

	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex-2] {
			nums[uniqueIndex] = nums[i]
			uniqueIndex++
		}
	}
	return uniqueIndex
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	result := removeDuplicates(nums)

	fmt.Printf("Unique elements: %d\n", result)
	fmt.Println("Unique array:", nums[:result])
}
