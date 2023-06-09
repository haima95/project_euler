package main

import (
	"fmt"
	"math"
)

/*   Triangular, pentagonal, and hexagonal

Triangle, pentagonal, and hexagonal numbers are generated by the following formulae:

Triangle	 	Tn=n(n+1)/2	 	1, 3, 6, 10, 15, ...
Pentagonal	 	Pn=n(3n−1)/2	 	1, 5, 12, 22, 35, ...
Hexagonal	 	Hn=n(2n−1)	 	1, 6, 15, 28, 45, ...
It can be verified that T285 = P165 = H143 = 40755.

Find the next triangle number that is also pentagonal and hexagonal.

*/

func triangularPentagonalAndHexagonal() int {
	t1 := 286
	for {
		k := t1 * (t1 + 1) / 2
		np, ok := getPentagonal(k)
		if ok {
			nh, ok := getHexagonal(k)
			if ok {
				fmt.Println(k, " : ", t1, ", ", np, ", ", nh)
				return k
			}
		}
		t1++
	}
	return -1
}

func getPentagonal(p int) (int, bool) {
	/*
		n*(3*n-1)/2 = p
		3*n*n - n = 2*p
		9*n*n - 3*n = 6*p
		(3*n - 1/2)^2 = 6*p + 1/4
		(6*n - 1)^2 = 24*p+1
	*/
	k := math.Sqrt(24*float64(p) + 1)
	s := int(k)
	// 如果不能开平方根，就不满足条件
	if s*s != 24*p+1 {
		return -1, false
	}
	return (s + 1) / 6, (s+1)%6 == 0
}

func getHexagonal(h int) (int, bool) {
	/*
		2n*n-n = h
		(4n-1)^2 = 8h+1
	*/

	k := math.Sqrt(8*float64(h) + 1)
	p := int(k)
	// 如果不能开平方根，就不满足条件
	if p*p != 8*h+1 {
		return -1, false
	}
	return (p + 1) / 4, (p+1)%4 == 0
}
