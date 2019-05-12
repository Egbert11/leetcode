// Source: https://leetcode.com/problems/valid-anagram/
// Author: Egbert11
// Date: 2019-05-12

/*
Share
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

	Input: s = "anagram", t = "nagaram"
	Output: true
Example 2:

	Input: s = "rat", t = "car"
	Output: false
Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?
*/

package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	seen := make([]int, 26)
	for _, c := range s {
		seen[c - 'a']++
	}
	for _, c := range t {
		seen[c - 'a']--
	}
	for _, i := range seen {
		if i != 0 {
			return false
		}
	}
	return true
}


func main() {
	s, t := "anagram", "nagaram"
	fmt.Println(isAnagram(s, t))
	s, t = "rat", "car"
	fmt.Println(isAnagram(s, t))
}
