// Source: https://leetcode.com/problems/word-pattern/
// Author: Egbert11
// Date: 2019-05-18

/*
Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.

Example 1:

	Input: pattern = "abba", str = "dog cat cat dog"
	Output: true
Example 2:

	Input:pattern = "abba", str = "dog cat cat fish"
	Output: false
Example 3:

	Input: pattern = "aaaa", str = "dog cat cat dog"
	Output: false
Example 4:

	Input: pattern = "abba", str = "dog dog dog dog"
	Output: false
Notes:
You may assume pattern contains only lowercase letters, and str contains lowercase letters that may be separated by a single space.
*/

package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	n := len(pattern)
	strs := strings.Split(str, " ")
	if len(strs) != n {
		return false
	}
	m1 := make(map[uint8]string)
	m2 := make(map[string]uint8)
	for i := 0; i < n; i++ {
		if mStr, ok := m1[pattern[i]]; ok {
			if mStr != strs[i] {
				return false
			}
		} else {
			m1[pattern[i]] = strs[i]
		}

		if c, ok := m2[strs[i]]; ok {
			if c != pattern[i] {
				return false
			}
		} else {
			m2[strs[i]] = pattern[i]
		}
	}
	return true
}

func main() {
	pattern := "abba"
	str := "dog cat cat dog"
	fmt.Println(wordPattern(pattern, str))

	pattern = "abba"
	str = "dog cat cat fish"
	fmt.Println(wordPattern(pattern, str))

	pattern = "aaaa"
	str = "dog cat cat dog"
	fmt.Println(wordPattern(pattern, str))

	pattern = "abba"
	str = "dog dog dog dog"
	fmt.Println(wordPattern(pattern, str))
}
