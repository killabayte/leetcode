/*
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Constraints:
  1 <= strs.length <= 200
  0 <= strs[i].length <= 200
  strs[i] consists of only lower-case English letters.
*/

package main

import (
    "fmt"
    "sort"
)

func main(strs []string) string {
    var result string

    if len(strs) < 0 {
        return ""
    } else if len(strs) == 1 {
        return (string(strs[0]))
    } else {
        sort.Strings(strs)
        for i := range strs[0] {
            if strs[0][i] == strs[len(strs)-1][i] {
                result += string(strs[0][i])
            } else {
                break
            }
        }
    }
    return result
}
