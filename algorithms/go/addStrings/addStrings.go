// Source: https://leetcode.com/problems/add-strings/
// Author: Egbert11
// Date: 2019-05-24

/*
Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

Note:

The length of both num1 and num2 is < 5100.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addStrings(num1 string, num2 string) string {
	var sum []int
	len1 := len(num1)
	len2 := len(num2)
	if len1 > len2 {
		len1, len2 = len2, len1
		num1, num2 = num2, num1
	}
	carrier := 0
	for i := 0; i < len1; i++ {
		c1, _ := strconv.Atoi(string(num1[len1-1-i]))
		c2, _ := strconv.Atoi(string(num2[len2-1-i]))
		s := c1 + c2 + carrier
		sum = append(sum, s%10)
		carrier = s/10
	}
	for i := len1; i < len2; i++ {
		s, _ := strconv.Atoi(string(num2[len2-1-i]))
		s += carrier
		sum = append(sum, s%10)
		carrier = s/10
	}
	if carrier == 1 {
		sum  = append(sum, carrier)
	}
	// reverse
	for i, j := 0, len(sum)-1; i < j; i, j = i+1, j-1 {
		sum[i], sum[j] = sum[j], sum[i]
	}
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sum)),""),"[]")
}

func addStrings1(num1 string, num2 string) string {
	c := byte(0)
	i1 := len(num1) - 1
	i2 := len(num2) - 1
	var sum []byte
	if i1 > i2 {
		sum = make([]byte, 0, i1+1)
	} else {
		sum = make([]byte, 0, i2+1)
	}
	for tmp := c; i1 >= 0 || i2 >= 0; tmp = c{
		if i1 >= 0 {
			tmp += num1[i1] - '0'
			i1 -= 1
		}
		if i2 >= 0 {
			tmp += num2[i2] - '0'
			i2 -= 1
		}
		if tmp >= 10 {
			c = 1
			tmp = tmp % 10
		} else {
			c = 0
		}
		sum = append(sum, tmp + '0')
	}
	if c == 1 {
		sum = append(sum, '1')
	}
	// reverse
	for i, j := 0, len(sum)-1; i < j; i, j = i+1, j-1 {
		sum[i], sum[j] = sum[j], sum[i]
	}
	return string(sum)
}


func main() {
	fmt.Println(addStrings("123141431431","32141341"))
	fmt.Println(addStrings1("123141431431","32141341"))
	fmt.Println(addStrings("1234","3238"))
	fmt.Println(addStrings1("1234","3238"))
	fmt.Println(addStrings("1","9"))
	fmt.Println(addStrings1("1","9"))
}
