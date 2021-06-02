/*
Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
Return the running sum of nums.

Example:
Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
*/

package main

func main(nums []int) []int {
    var sumSlice []int
    var previousSum int
    
    for i := range nums {
        if i == 0 {
            sumSlice = append(sumSlice, nums[i])
            previousSum += nums[i]
        } else if i < (len(nums)) && i != 0 {
            sumSlice = append(sumSlice, nums[i]+previousSum)
            previousSum += nums[i]
        }
    }
    return sumSlice
}
