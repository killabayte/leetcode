package main

import "fmt"

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))

	for _, i := range nums[1:] {
		ans = append(ans, ans[-1]*nums[i-1])
	}
	fmt.Println(ans)
	return ans
}
