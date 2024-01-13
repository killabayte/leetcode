package main

func removeElement(nums []int, val int) int {
	i := 0
	for j, num := range nums {
		if num != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
