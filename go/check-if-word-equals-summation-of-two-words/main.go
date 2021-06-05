/*
The letter value of a letter is its position in the alphabet starting from 0 (i.e. 'a' -> 0, 'b' -> 1, 'c' -> 2, etc.).

The numerical value of some string of lowercase English letters s is the concatenation of the letter values of each letter in s, which is then converted into an integer.

For example, if s = "acb", we concatenate each letter's letter value, resulting in "021". After converting it, we get 21.
You are given three strings firstWord, secondWord, and targetWord, each consisting of lowercase English letters 'a' through 'j' inclusive.

Return true if the summation of the numerical values of firstWord and secondWord equals the numerical value of targetWord, or false otherwise.

Example 1:
Input: firstWord = "acb", secondWord = "cba", targetWord = "cdb"
Output: true
Explanation:
The numerical value of firstWord is "acb" -> "021" -> 21.
The numerical value of secondWord is "cba" -> "210" -> 210.
The numerical value of targetWord is "cdb" -> "231" -> 231.
We return true because 21 + 210 == 231.

Example 2:
Input: firstWord = "aaa", secondWord = "a", targetWord = "aab"
Output: false
Explanation:
The numerical value of firstWord is "aaa" -> "000" -> 0.
The numerical value of secondWord is "a" -> "0" -> 0.
The numerical value of targetWord is "aab" -> "001" -> 1.
We return false because 0 + 0 != 1.
*/

package main

import (
	"fmt"
	"strconv"
)

func main(firstWord string, secondWord string, targetWord string) bool {
	var firstWordSum int
	var secondWordSum int
	var targetWordSum int

	firstWordSum = wordSumMain(firstWord)
	secondWordSum = wordSumMain(secondWord)
	targetWordSum = wordSumMain(targetWord)
	fmt.Println(firstWordSum)
	fmt.Println(secondWordSum)
	fmt.Println(targetWordSum)

	return (firstWordSum+secondWordSum == targetWordSum)
}

func wordSumMain(word string) int {
	var wordSum string

	for _, v := range word {
		if string(v) == "a" {
			wordSum += string("0")
		} else if string(v) == "b" {
			wordSum += string("1")
		} else if string(v) == "c" {
			wordSum += string("2")
		} else if string(v) == "d" {
			wordSum += string("3")
		} else if string(v) == "e" {
			wordSum += string("4")
		} else if string(v) == "f" {
			wordSum += string("5")
		} else if string(v) == "g" {
			wordSum += string("6")
		} else if string(v) == "h" {
			wordSum += string("7")
		} else if string(v) == "i" {
			wordSum += string("8")
		} else if string(v) == "j" {
			wordSum += string("9")
		}
	}
	wordSumResult, _ := strconv.Atoi(wordSum)
	return wordSumResult
}
