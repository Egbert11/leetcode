// Source: https://leetcode.com/problems/sum-of-two-integers/
// Author: Egbert11
// Date: 2019-05-18

/*
Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.

Example 1:

	Input: a = 1, b = 2
	Output: 3
Example 2:

	Input: a = -2, b = 3
	Output: 1
*/

package main

import "fmt"

func getSum(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return getSum(a^b, (a&b)<< 1) //be careful about the terminating condition;
	}
}

func main() {
	fmt.Println(getSum(-5, 10))
	fmt.Println(getSum(2, 3))
}
