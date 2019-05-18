// Source: https://leetcode.com/problems/range-sum-query-immutable/
// Author: Egbert11
// Date: 2019-05-18

/*
Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

Example:
	Given nums = [-2, 0, 3, -5, 2, -1]

	sumRange(0, 2) -> 1
	sumRange(2, 5) -> -1
	sumRange(0, 5) -> -3
Note:
	1.You may assume that the array does not change.
	2.There are many calls to sumRange function.
*/

package main

import "fmt"

type NumArray struct {
	sum []int
}


func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums)+1)

	for i, num := range nums {
		sum[i+1] = sum[i] + num
	}

	return NumArray{sum}
}


func (this *NumArray) SumRange(i int, j int) int {
	return this.sum[j+1] - this.sum[i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func main() {
	obj := Constructor([]int{1,2,3,4,5,6,7})
	fmt.Println(obj.SumRange(1,2))
	fmt.Println(obj.SumRange(0,5))
}
