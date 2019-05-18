// Source: https://leetcode.com/problems/missing-number/
// Author: Egbert11
// Date: 2019-05-18

/*
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example 1:

	Input: [3,0,1]
	Output: 2
Example 2:

	Input: [9,6,4,2,3,5,7,0,1]
	Output: 8
Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?
*/

package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	total := n*(n+1)/2
	for _, x := range nums {
		total -= x
	}
	return total
}

func main() {
	nums := []int{9,6,4,2,3,5,7,0,1}
	fmt.Println(missingNumber(nums))
	nums = []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}
