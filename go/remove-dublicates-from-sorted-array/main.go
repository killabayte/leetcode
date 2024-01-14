package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	uniqueIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex] {
			uniqueIndex++
			nums[uniqueIndex] = nums[i]
		}
	}

	return uniqueIndex + 1
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5}
	result := removeDuplicates(nums)

	fmt.Printf("Unique elements: %d\n", result)
	fmt.Println("Unique array:", nums[:result])
}
