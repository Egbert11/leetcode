// Source: https://leetcode.com/problems/two-sum/
// Author: Egbert11
// Date: 2019-05-06

/*
	Given an array of integers, return indices of the two numbers such that they add up to a specific target.

	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	Example:

	Given nums = [2, 7, 11, 15], target = 9,

	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].
*/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indice_map := make(map[int]int)
	for i, num := range nums {
		indice_map[num] = i
	}
	var res []int
	for i, num := range nums {
		other := target - num
		if indice, ok := indice_map[other]; ok {
			if indice != i {
				res = append(res, i)
				res = append(res, indice)
				break
			}
		}
	}
	return res
}

func main()  {
	nums := []int{2,7,11,15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
	nums = []int{3, 2, 4}
	target = 6
	res = twoSum(nums, target)
	fmt.Println(res)
}
