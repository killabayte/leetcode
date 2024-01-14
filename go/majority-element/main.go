package main

import "fmt"

func majorityElement(nums []int) int {
	majorityCount := len(nums) / 2

	recorded := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		recorded[num]++

		if count := recorded[num]; count > majorityCount {
			return num
		}
	}

	return -1
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 1, 1, 2, 2}
	result := majorityElement(nums)

	fmt.Printf("Majority element: %d\n", result)
}
