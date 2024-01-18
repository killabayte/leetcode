package main

import "fmt"

func rotate(nums []int, k int) {
	rotated_nums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		rotated_nums[(i+k)%len(nums)] = nums[i]
	}
	nums = rotated_nums
	fmt.Println(nums)
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
