// Source: https://leetcode.com/problems/excel-sheet-column-title/
// Author: Egbert11
// Date: 2019-05-07

/*
	Given a positive integer, return its corresponding column title as appear in an Excel sheet.

	For example:

		1 -> A
		2 -> B
		3 -> C
		...
		26 -> Z
		27 -> AA
		28 -> AB
		...
	Example 1:
		Input: 1
		Output: "A"

	Example 2:
		Input: 28
		Output: "AB"

	Example 3:
		Input: 701
		Output: "ZY"
*/

package main

import (
	"bytes"
	"fmt"
)

func convertToTitle(n int) string {
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var buffer bytes.Buffer
	for n > 0 {
		quotient := n / 26
		reminder := n % 26
		if reminder == 0 {
			buffer.WriteByte(alpha[25])
			quotient -= 1
		} else {
			buffer.WriteByte(alpha[reminder-1])
		}
		n = quotient
	}
	return reverse(buffer.String())
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main()  {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(2))
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(27))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(52))
	fmt.Println(convertToTitle(53))
	fmt.Println(convertToTitle(701))
}
