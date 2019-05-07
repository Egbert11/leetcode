// Source: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// Author: Egbert11
// Date: 2019-05-07

/*
	Given an array of integers that is already sorted in ascending order,
	find two numbers such that they add up to a specific target number.

	The function twoSum should return indices of the two numbers such that they add up to the
	target, where index1 must be less than index2.

	Note:

		Your returned answers (both index1 and index2) are not zero-based.
		You may assume that each input would have exactly one solution and you may not use the
		same element twice.

	Example:

		Input: numbers = [2,7,11,15], target = 9
		Output: [1,2]
		Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
*/

package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	var begin = 0
	var end = len(numbers) - 1
	res := []int{}
	for begin < end {
		total := numbers[begin] + numbers[end]
		if total == target {
			res = append(res, begin+1)
			res = append(res, end+1)
			break
		} else if total < target {
			begin += 1
		} else {
			end -= 1
		}
	}
	return res

}

func main()  {
	numbers := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}
