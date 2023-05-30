package main

import "fmt"

/*   Spiral Primes

Starting with 1 and spiralling anticlockwise in the following way, a square spiral with side length 7 is formed.

37 36 35 34 33 32 31
38 17 16 15 14 13 30
39 18  5  4  3 12 29
40 19  6  1  2 11 28
41 20  7  8  9 10 27
42 21 22 23 24 25 26
43 44 45 46 47 48 49

It is interesting to note that the odd squares lie along the bottom right diagonal, but what is more interesting is that
8 out of the 13 numbers lying along both diagonals are prime; that is, a ratio of 8/13 ≈ 62%.

If one complete new layer is wrapped around the spiral above, a square spiral with side length 9 will be formed. If this
process is continued, what is the side length of the square spiral for which the ratio of primes along both diagonals
first falls below 10%?

*/

func spiralPrimes() int {
	primes := []int{3, 5, 7}
	values := []int{1, 3, 5, 7, 9}
	n := 5
	for {
		d := values[len(values)-1]
		for i := 1; i < 5; i++ {
			u := d + i*(n-1)
			values = append(values, u)
			if isPrime(u) {
				primes = append(primes, u)
			}
		}
		if float64(len(primes))/float64(len(values)) < 0.1 {
			break
		}
		if n < 14 {
			fmt.Printf(" --- %d --- \n", n)
			fmt.Println(primes)
			fmt.Println(values)
			fmt.Printf(" --- --- --- \n")
		}
		n += 2
	}
	return n
}
