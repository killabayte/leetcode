package main

func subarraySum(nums []int, k int) int {
	answer := 0
	sum := 0
	prefixSumCount := map[int]int{0: 1}

	for _, num := range nums {
		sum += num
		toRemove := sum - k
		answer += prefixSumCount[toRemove]
		prevCount := prefixSumCount[sum]
		prefixSumCount[sum] = prevCount + 1
	}
	return answer
}

func main() {
	exampleOne := subarraySum([]int{1, 1, 1}, 2)
	println(exampleOne)

}
