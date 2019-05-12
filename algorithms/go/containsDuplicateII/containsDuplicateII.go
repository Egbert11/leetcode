// Source: https://leetcode.com/problems/contains-duplicate-ii/
// Author: Egbert11
// Date: 2019-05-12

/*
Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array
such that nums[i] = nums[j] and the absolute difference between i and j is at most k.

Example 1:

	Input: nums = [1,2,3,1], k = 3
	Output: true
Example 2:

	Input: nums = [1,0,1,1], k = 1
	Output: true
Example 3:

	Input: nums = [1,2,3,1,2,3], k = 2
	Output: false
*/

package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		if _, ok := m[num]; ok {
			if i - m[num] <= k {
				return true
			}
		}
		m[num] = i
	}
	return false
}


func main() {
	nums := []int{1,2,3,1}
	fmt.Println(containsNearbyDuplicate(nums, 3))
	nums = []int{1,0,1,1}
	fmt.Println(containsNearbyDuplicate(nums, 1))
	nums = []int{1,2,3,1,2,3}
	fmt.Println(containsNearbyDuplicate(nums, 2))
}
