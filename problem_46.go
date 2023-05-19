package main

import (
	"fmt"
	"math"
)

/*    Goldbach's other conjecture

It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.

9 = 7 + 2×12
15 = 7 + 2×22
21 = 3 + 2×32
25 = 7 + 2×32
27 = 19 + 2×22
33 = 31 + 2×12

It turns out that the conjecture was false.

What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?

*/

func goldbachsOtherConjecture() int {
	d := 3
	primes := []int{2}
	for {
		if isPrime(d) {
			primes = append(primes, d)
			d += 2
			continue
		}
		i := len(primes) - 1
		for ; i >= 0; i-- {
			k := d - primes[i]
			q := math.Sqrt(float64(k / 2))
			if int(q)*int(q) == k/2 {
				fmt.Println(d, ":", primes[i], "+2*", q, "*", q)
				break
			}
		}
		if i == -1 {
			return d
		}
		d += 2
		//if d > 40 {
		//	break
		//}
	}
	return -1
}
