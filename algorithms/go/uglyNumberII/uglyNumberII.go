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

/*
Explanation:

The key is to realize each number can be and have to be generated by a former number multiplied by 2, 3 or 5
e.g.
1 2 3 4 5 6 8 9 10 12 15..
what is next?
it must be x * 2 or y * 3 or z * 5, where x, y, z is an existing number.

How do we determine x, y, z then?
apparently, you can just traverse the sequence generated by far from 1 ... 15, until you find such x, y, z that x * 2,
y * 3, z * 5 is just bigger than 15. In this case x=8, y=6, z=4. Then you compare x * 2, y * 3, z * 5 so you know
next number will be x * 2 = 8 * 2 = 16.
k, now you have 1,2,3,4,....,15, 16,

Then what is next?
You wanna do the same process again to find the new x, y, z, but you realize, wait, do I have to
traverse the sequence generated by far again?

NO! since you know last time, x=8, y=6, z=4 and x=8 was used to generate 16, so this time, you can immediately know
the new_x = 9 (the next number after 8 is 9 in the generated sequence), y=6, z=4.
Then you need to compare new_x * 2, y * 3, z * 5. You know next number is 9 * 2 = 18;

And you also know, the next x will be 10 since new_x = 9 was used this time.
But what is next y? apparently, if y=6, 6*3 = 18, which is already generated in this round. So you also need to
update next y from 6 to 8.

Based on the idea above, you can actually generated x,y,z from very beginning, and update x, y, z accordingly.
It ends up with a O(n) solution.
*/
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
