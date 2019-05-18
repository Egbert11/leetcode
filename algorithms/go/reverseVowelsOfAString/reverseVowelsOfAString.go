// Source: https://leetcode.com/problems/reverse-vowels-of-a-string/
// Author: Egbert11
// Date: 2019-05-18

/*
Write a function that takes a string as input and reverse only the vowels of a string.

Example 1:

	Input: "hello"
	Output: "holle"
Example 2:

	Input: "leetcode"
	Output: "leotcede"
Note:
The vowels does not include the letter "y".
*/

package main

import (
	"fmt"
)

func reverseVowels(s string) string {
	sArr := []byte(s)
	vowels := map[byte]bool{
		'a':true,'e':true,'i':true,'o':true,'u':true,
		'A':true,'E':true,'I':true,'O':true,'U':true,
		}
	var indexes []int
	for i, c := range sArr {
		if _, ok := vowels[byte(c)]; ok {
			indexes = append(indexes, i)
		}
	}
	n := len(indexes)
	for i := 0; i < n/2; i++ {
		sArr[indexes[i]], sArr[indexes[n-1-i]] = sArr[indexes[n-1-i]], sArr[indexes[i]]
	}
	return string(sArr)
}

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
}
