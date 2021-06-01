package main

func main(nums []int, target int) []int {

	var result []int
main:
	for i := range nums {
		for j := range nums {
			if i < j && nums[i]+nums[j] == target {
				result = append(result, i, j)
				break main
			}
		}
	}
	return result
}
