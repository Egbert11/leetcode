// Source: https://leetcode.com/problems/ugly-number/
// Author: Egbert11
// Date: 2019-05-12

/*
Write a program to check whether a given number is an ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.

Example 1:

	Input: 6
	Output: true
	Explanation: 6 = 2 × 3
Example 2:

	Input: 8
	Output: true
	Explanation: 8 = 2 × 2 × 2
Example 3:

	Input: 14
	Output: false
	Explanation: 14 is not ugly since it includes another prime factor 7.
Note:

1 is typically treated as an ugly number.
Input is within the 32-bit signed integer range: [−231,  231 − 1].
*/

package main

import "fmt"

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	if num == 1 {
		return true
	}
	for num != 1 {
		if num % 2 == 0 {
			num /= 2
			continue
		}
		if num % 3 == 0 {
			num /= 3
			continue
		}
		if num % 5 == 0 {
			num /= 5
			continue
		}
		return false
	}
	return true
}


func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(isUgly(i))
	}
}
