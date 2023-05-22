package main

import (
	"math"
)

/*  Pentagon numbers

Pentagonal numbers are generated by the formula, Pn=n(3n−1)/2. The first ten pentagonal numbers are:

1, 5, 12, 22, 35, 51, 70, 92, 117, 145, ...

It can be seen that P4 + P7 = 22 + 70 = 92 = P8. However, their difference, 70 − 22 = 48, is not pentagonal.

Find the pair of pentagonal numbers, Pj and Pk, for which their sum and difference are pentagonal and D = |Pk − Pj| is
minimised; what is the value of D?

*/

// 参考 https://www.educative.io/answers/project-euler-44-pentagonal-numbers
func pentagonNumbers() int {
	i := 0
	for {
		i += 1
		k := i * (3*i - 1) / 2
		for v := 1; v < i; v++ {
			j := v * (3*v - 1) / 2
			if checkValue(k-j) && checkValue(k+j) {
				return k - j
			}
		}
	}
	return -1
}

func checkValue(n int) bool {
	k := math.Sqrt(1 + 24*float64(n))
	if int(k)*int(k) != 1+24*n {
		return false
	}
	return int(1+k)%6 == 0
}