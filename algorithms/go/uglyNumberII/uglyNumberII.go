// Source: https://leetcode.com/problems/ugly-number-ii/
// Author: Egbert11
// Date: 2019-05-12

/*
Write a program to find the n-th ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.

Example:

	Input: n = 10
	Output: 12
	Explanation: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.
Note:

1 is typically treated as an ugly number.
n does not exceed 1690.
*/

package main

import (
	"fmt"
	"math"
)

//dp
func nthUglyNumber(n int) int {
	arr := make([]float64, n, n)
	arr[0] = 1
	t2, t3, t5 := 0, 0, 0
	for i := 1; i < n; i++ {
		arr[i] = math.Min(arr[t2]*2, math.Min(arr[t3]*3, arr[t5]*5))
		if arr[i] == arr[t2]*2 {
			t2++
		}
		if arr[i] == arr[t3]*3 {
			t3++
		}
		if arr[i] == arr[t5]*5 {
			t5++
		}
	}
	return int(arr[n-1])
}


func main() {
	fmt.Println(nthUglyNumber(10))
	fmt.Println(nthUglyNumber(100))
	fmt.Println(nthUglyNumber(1690))
}
