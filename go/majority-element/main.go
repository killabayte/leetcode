package main

import "fmt"

func majorityElement(nums []int) int {
	numElement := 0
	middleEntry := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			numElement = nums[i]
			middleEntry++
		} else {
			if nums[i] == numElement {
				middleEntry++
			}
		}
	}
	return middleEntry
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	result := majorityElement(nums)

	fmt.Printf("Majority element: %d\n", result)
}
