// Source: https://leetcode.com/problems/power-of-three/
// Author: Egbert11
// Date: 2019-05-18

/*
Given an integer, write a function to determine if it is a power of three.

Example 1:

	Input: 27
	Output: true
Example 2:

	Input: 0
	Output: false
Example 3:

	Input: 9
	Output: true
Example 4:

	Input: 45
	Output: false
Follow up:
Could you do it without using any loop / recursion?
*/

package main

import "fmt"

func isPowerOfThree(n int) bool {
	max3PowerInt := 1162261467 // 3^19, 3^20 = 3486784401 > MaxInt32
	if n <= 0 || n > max3PowerInt {
		return false
	} else {
		return max3PowerInt % n == 0
	}
}

func main() {
	fmt.Println(isPowerOfThree(3))
	fmt.Println(isPowerOfThree(81))
	fmt.Println(isPowerOfThree(45))
}
