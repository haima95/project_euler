package main

import (
	"fmt"
	"strconv"
)

/*   Champernowne's constant

An irrational decimal fraction is created by concatenating the positive integers:

0.123456789101112131415161718192021...

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the following expression.

d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000

*/

func champernownesConstant() int {
	i := 1
	s := "0"
	for len(s) < 1000001 {
		d := strconv.Itoa(i)
		i++
		s += d
	}
	ans := 1
	arr := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	fmt.Println(s[:100])
	for _, v := range arr {
		fmt.Printf("%c ", s[v])
		ans *= int(s[v] - '0')
	}
	fmt.Println()
	return ans
}
