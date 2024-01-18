package main

func rotate(nums []int, k int) {
	n := len(nums)

	// Handle cases where k is greater than the array length
	k %= n

	// Reverse the entire array
	reverse(nums, 0, n-1)

	// Reverse the first k elements
	reverse(nums, 0, k-1)

	// Reverse the remaining elements
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
