// Source: https://leetcode.com/problems/intersection-of-two-arrays/
// Author: Egbert11
// Date: 2019-05-18

/*
Given two arrays, write a function to compute their intersection.

Example 1:

	Input: nums1 = [1,2,2,1], nums2 = [2,2]
	Output: [2]
Example 2:

	Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
	Output: [9,4]
Note:

Each element in the result must be unique.
The result can be in any order.
*/

package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	var result []int
	if n1 == 0 || n2 == 0 {
		return result
	}
	if n1 > n2 {
		n1, n2 = n2, n1
		nums1, nums2 = nums2, nums1
	}
	var m1  = make(map[int]bool)
	var m2  = make(map[int]bool)
	for _, n := range nums1 {
		m1[n] = true
	}
	for _, n := range nums2 {
		if _, ok := m1[n]; ok {
			m2[n] = true
		}
	}
	for k, _ := range m2 {
		result = append(result, k)
	}
	return result
}

func main() {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	fmt.Println(intersection(nums1, nums2))
	nums1 = []int{4,9,5}
	nums2 = []int{9,4,9,8,4}
	fmt.Println(intersection(nums1, nums2))
}
