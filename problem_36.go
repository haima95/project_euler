package main

import "fmt"

/*    Double-base palindromes

The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)

*/

func double_basePalindromes() int {
	sum := 0
	for i := 1; i < 1000000; i += 2 {
		ok10, base10 := divBase(i, 10)
		if !ok10 {
			continue
		}
		ok2, base2 := divBase(i, 2)
		if ok2 {
			sum += i
			fmt.Print(i, " : ")
			for _, v := range base10 {
				fmt.Print(v)
			}
			fmt.Print(" , ")
			for _, v := range base2 {
				fmt.Print(v)
			}
			fmt.Println()
		}
	}
	return sum
}

func divBase(n int, base int) (bool, []int) {
	var tmp []int
	for n > 0 {
		tmp = append(tmp, n%base)
		n /= base
	}
	m := len(tmp)
	for i := 0; i < m/2; i++ {
		if tmp[i] != tmp[m-i-1] {
			return false, []int{}
		}
	}
	return true, tmp
}
