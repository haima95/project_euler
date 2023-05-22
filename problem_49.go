package main

import (
	"fmt"
	"strconv"
)

/*   Prime permutations

The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual in two ways: (i)
each of the three terms are prime, and, (ii) each of the 4-digit numbers are permutations of one another.

There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting this property, but there is
one other 4-digit increasing sequence.

What 12-digit number do you form by concatenating the three terms in this sequence?

*/

func primePermutations() []string {
	primes := []int{}
	primeMap := make(map[int]bool)
	for i := 1001; i < 10000; i++ {
		if isPrime(i) {
			primes = append(primes, i)
			primeMap[i] = true
		}
	}
	fmt.Println(len(primes))
	ans := make([][3]int, 0)
	for i := 0; i < len(primes); i++ {
		for j := i + 1; j < len(primes); j++ {
			if checkPrime(primes[j], primes[i]) {
				d := primes[j] - primes[i]
				if primeMap[primes[j]+d] && checkPrime(primes[j]+d, primes[j]) {
					ans = append(ans, [3]int{primes[i], primes[j], primes[j] + d})
				}
			}

		}
	}
	fmt.Println(ans)
	s := make([]string, len(ans))
	for i, v := range ans {
		s[i] = ""
		for _, v1 := range v {
			ss := strconv.Itoa(v1)
			s[i] += ss
		}

	}
	return s
}

func checkPrime(n, m int) bool {
	tmp1 := make(map[int]int)
	tmp2 := make(map[int]int)
	for n > 0 {
		tmp1[n%10]++
		n /= 10
	}
	for m > 0 {
		tmp2[m%10]++
		m /= 10
	}
	if len(tmp1) != len(tmp2) {
		return false
	}
	for k, v := range tmp1 {
		if tmp2[k] != v {
			return false
		}
	}
	return true
}
