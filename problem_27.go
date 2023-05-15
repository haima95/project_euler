package main

import "fmt"

/*	Quadratic primes

 */

func quadraticPrimes() int {
	max, value := 0, 0
	for a := 999; a > -1000; a-- {
		for b := 1000; b > -1001; b-- {
			n := 0
			count := 0
			for {
				k := n*n + n*a + b
				if isPrime(k) {
					count++
				} else {
					//if a == 1 && b == 41 {
					//	fmt.Println(a, b, ": ", count)
					//}
					//if a == -79 && b == 1601 {
					//	fmt.Println(a, b, ": ", count)
					//}
					if count > max {
						max = count
						value = a * b
						fmt.Println(a, b, ": ", count, " - ", value)
					}
					break
				}
				n++
			}
		}
	}
	return value
}
