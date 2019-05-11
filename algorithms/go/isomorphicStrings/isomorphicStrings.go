// Source: https://leetcode.com/problems/isomorphic-strings/
// Author: Egbert11
// Date: 2019-05-12

/*
Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters.
No two characters may map to the same character but a character may map to itself.

Example 1:

	Input: s = "egg", t = "add"
	Output: true
Example 2:

	Input: s = "foo", t = "bar"
	Output: false
Example 3:

	Input: s = "paper", t = "title"
	Output: true
Note:
You may assume both s and t have the same length.
*/

package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	return check(s, t) && check(t, s)
}

func check(s string, t string) bool {
	m := make(map[int32]uint8)
	for i, c := range s {
		tC := t[i]
		if v, ok := m[c]; ok {
			if v != tC {
				return false
			}
		} else {
			m[c] = tC
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
}
