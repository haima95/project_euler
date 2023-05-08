package main

import "fmt"

/*		Digit cancelling fractions

The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.

*/

func digitCancellingFractions() int {
	tmp := make(map[string][2]int)
	// We shall consider fractions like, 30/50 = 3/5, to be trivial examples.
	// 因为求的是no-trivial，所以去掉i = 0
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			t1 := j*10 + i
			for k := 0; k < 10; k++ {
				t2 := i*10 + k
				if t1 < t2 && t1*k == t2*j {
					fmt.Println(t1, "/", t2, " == ", j, "/", k)
					tmp[fmt.Sprintf("%d_%d", j, k)] = [2]int{j, k}
				}
			}
		}
	}

	// 四个分数相乘得到一个分数，将分数的分子与分母除以公约数，得到最小常用项
	numerator := 1
	denominator := 1
	for _, v := range tmp {
		numerator *= v[0]
		denominator *= v[1]
	}
	k := gcd(numerator, denominator)
	return denominator / k
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
