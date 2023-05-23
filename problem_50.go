package main

import "fmt"

/*    Consecutive prime sum

The prime 41, can be written as the sum of six consecutive primes:

41 = 2 + 3 + 5 + 7 + 11 + 13
This is the longest sum of consecutive primes that adds to a prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most consecutive primes?

*/

func consecutivePrimeSum(n int) int {
	primes := []int{2}
	primeM := map[int]bool{2: true}
	max, index := 1, 0
	for i := 3; i <= n; i += 2 {
		if isPrime(i) {
			primes = append(primes, i)
			primeM[i] = true
		}
	}
	for i := 0; i < len(primes); i++ {
		sum := primes[i]
		for j := i + 1; j < len(primes); j++ {
			sum += primes[j]
			if sum > n {
				break
			}
			if primeM[sum] {
				if j+1-i > max {
					max = j + 1 - i
					index = i
				}
			}
		}
	}
	sum := 0
	for i := 0; i < max; i++ {
		sum += primes[i+index]
		fmt.Print(primes[index+i], " ")
	}
	fmt.Println()
	return sum
}
