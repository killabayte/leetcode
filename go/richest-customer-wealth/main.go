/*
You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank. Return the wealth that the richest customer has.

A customer's wealth is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum wealth.

Example:
Input: accounts = [[1,2,3],[3,2,1]]
Output: 6
Explanation:
1st customer has wealth = 1 + 2 + 3 = 6
2nd customer has wealth = 3 + 2 + 1 = 6
Both customers are considered the richest with a wealth of 6 each, so return 6.
*/

package main

func main(accounts [][]int) int {
	var maxWealth int

	for i := range accounts {
		var currentWealth int

		for _, jv := range accounts[i] {
			currentWealth += jv
		}
		if currentWealth == maxWealth {
			continue
		} else if currentWealth > maxWealth {
			maxWealth = currentWealth
		}
	}

	return maxWealth
}
