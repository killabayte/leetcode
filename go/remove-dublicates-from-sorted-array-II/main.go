package main

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
