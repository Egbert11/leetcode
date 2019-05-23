// Source: https://leetcode.com/problems/fizz-buzz/
// Author: Egbert11
// Date: 2019-05-23

/*
Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of
five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

Example:

n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
*/

package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	result := []string{}
	for i := 1; i <= n; i++ {
		if i % 3 != 0 && i % 5 != 0{
			result = append(result, strconv.Itoa(i))
		} else if i % 15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i % 3 == 0 {
			result = append(result, "Fizz")
		} else {
			result = append(result, "Buzz")
		}
	}
	return result
}


func main() {
	fmt.Println(fizzBuzz(15))
}
