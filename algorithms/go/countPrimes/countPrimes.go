// Source: https://leetcode.com/problems/count-primes/
// Author: Egbert11
// Date: 2019-05-11

/*
Count the number of prime numbers less than a non-negative number, n.

Example:

	Input: 10
	Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
*/

package main

import "fmt"

func countPrimes(n int) int {
	arr := make([]int, n, n)
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			arr[i] = 0
		} else {
			arr[i] = 1
		}
	}
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			continue
		} else {
			step := i
			for j := step*2; j < n; j += step {
				arr[j] = 0
			}
		}
	}
	count := 0
	for _, num := range arr {
		if num == 1 {
			count += 1
		}
	}
	return count
}

func countPrimes2(n int) int {
	if n < 3 {
		return 0
	}
	var c int = n / 2
	f := make([]bool, n, n) // false means prime
	for i := 3; i < n; i += 2 {
		if f[i] {
			continue
		}

		for j := i*i; j < n; j += 2*i {
			if !f[j] {
				c--
				f[j] = true
			}
		}
	}
	return c

}

func main() {
	fmt.Println(countPrimes(1))
	fmt.Println(countPrimes(2))
	fmt.Println(countPrimes(10))
	fmt.Println(countPrimes(20))
	fmt.Println(countPrimes(50))
	fmt.Println(countPrimes(100))
	fmt.Println(countPrimes(1000))

	fmt.Println(countPrimes2(1))
	fmt.Println(countPrimes2(2))
	fmt.Println(countPrimes2(10))
	fmt.Println(countPrimes2(20))
	fmt.Println(countPrimes2(50))
	fmt.Println(countPrimes2(100))
	fmt.Println(countPrimes2(1000))
}
