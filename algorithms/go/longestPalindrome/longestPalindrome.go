// Source: https://leetcode.com/problems/longest-palindrome/
// Author: Egbert11
// Date: 2019-05-23

/*
Given a string which consists of lowercase or uppercase letters, find the length of the longest
palindromes that can be built with those letters.

This is case sensitive, for example "Aa" is not considered a palindrome here.

Note:
Assume the length of given string will not exceed 1,010.

Example:

	Input:
	"abccccdd"

	Output:
	7

Explanation:
One longest palindrome that can be built is "dccaccd", whose length is 7.
*/

package main

import "fmt"

func longestPalindrome(s string) int {
	m := make(map[byte]int)
	for _, c := range s {
		if _, ok := m[byte(c)]; ok {
			m[byte(c)] += 1
		} else {
			m[byte(c)] = 1
		}
	}
	oddFlag := false
	sum := 0
	for _, v := range m {
		if v % 2 == 1 {
			oddFlag = true
			sum += v - 1
		} else {
			sum += v
		}
	}
	if oddFlag {
		sum += 1
	}
	return sum
}


func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}
