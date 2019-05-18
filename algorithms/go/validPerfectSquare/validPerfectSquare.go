// Source: https://leetcode.com/problems/valid-perfect-square/
// Author: Egbert11
// Date: 2019-05-18

/*
Given a positive integer num, write a function which returns True if num is a perfect square else False.

Note: Do not use any built-in library function such as sqrt.

Example 1:

	Input: 16
	Output: true
Example 2:

	Input: 14
	Output: false
*/

package main

import "fmt"

func isPerfectSquare(num int) bool {
	var low, high = 1, num
	for low <= high {
		medium := (low + high) / 2
		multiple := medium* medium
		if multiple == num {
			return true
		} else if multiple < num {
			low = medium + 1
		} else {
			high = medium - 1
		}
	}
	return false
}

func main() {
	for i := 1; i < 30; i++ {
		fmt.Println(isPerfectSquare(i))
	}
}
