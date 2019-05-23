// Source: https://leetcode.com/problems/third-maximum-number/
// Author: Egbert11
// Date: 2019-05-23

/*
Given a non-empty array of integers, return the third maximum number in this array. If it does not
exist, return the maximum number. The time complexity must be in O(n).

Example 1:
Input: [3, 2, 1]

Output: 1

Explanation: The third maximum is 1.
Example 2:
Input: [1, 2]

Output: 2

Explanation: The third maximum does not exist, so the maximum (2) is returned instead.
Example 3:
Input: [2, 2, 3, 1]

Output: 1

Explanation: Note that the third maximum here means the third maximum distinct number.
Both numbers with value 2 are both considered as second maximum.
*/

package main

import (
	"fmt"
)

func thirdMax(nums []int) int {
	n := len(nums)
	minValue := nums[0]
	for i := 1; i < n; i++ {
		if minValue > nums[i] {
			minValue = nums[i]
		}
	}
	arr := []int{minValue, minValue, minValue}
	for _, num := range nums {
		if num == arr[0] || num == arr[1] || num == arr[2] {
			continue
		}
		if num > arr[0] {
			arr[2] = arr[1]
			arr[1] = arr[0]
			arr[0] = num
		} else if num > arr[1] {
			arr[2] = arr[1]
			arr[1] = num
		} else if num > arr[2] {
			arr[2] = num
		}
	}
	if arr[0] == arr[1] || arr[1] == arr[2] {
		return arr[0]
	} else {
		return arr[2]
	}
}


func main() {
	fmt.Println(thirdMax([]int{3, 2, 1}))
	fmt.Println(thirdMax([]int{1, 2}))
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
}
