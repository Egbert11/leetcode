// Source: https://leetcode.com/problems/power-of-four/
// Author: Egbert11
// Date: 2019-05-18

/*
Given an integer (signed 32 bits), write a function to check whether it is a power of 4.

Example 1:

	Input: 16
	Output: true
Example 2:

	Input: 5
	Output: false
Follow up: Could you solve it without loops/recursion?
*/

package main

import "fmt"

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 || n == 4 {
		return true
	}
	for {
		if n % 4 == 0 {
			return isPowerOfFour(n/4)
		}
		return false
	}
}

func isPowerOfFour2(n int) bool {
	return n > 0 && (n&(n-1)) == 0 && (n&0x55555555) != 0
}

func main() {
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(12))
	fmt.Println(isPowerOfFour(64))
}
