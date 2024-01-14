package main

import "fmt"

func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}

		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}

	count = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		}
	}

	if count > len(nums)/2 {
		return candidate
	}

	return -1
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 1, 1, 2, 2}
	result := majorityElement(nums)

	fmt.Printf("Majority element: %d\n", result)
}
