package main

func subarraySum(nums []int, k int) int {
	answer := 0
	sum := 0
	prefixSumCount := map[int]int{0: 1}

	for _, i := range nums {
		sum += nums[i]
		toRemove := sum - k
		answer += prefixSumCount[toRemove]
		prevCount := prefixSumCount[sum]
		prefixSumCount[sum] = prevCount + 1
	}
}
