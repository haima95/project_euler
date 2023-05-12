package main

import "fmt"

/*    Truncatable primes

The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left
to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37,
and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.

*/

func truncatablePrimes() int {
	base := []int{1, 1}
	var arr []int
	for len(arr) != 11 {
		k := 0
		n := len(base)
		i := 0
		for ; i < n; i++ {
			k = k*10 + base[i]
			if !isPrime3(k) {
				break
			}
		}
		if i == n {
			k1 := 0
			i = n - 1
			m := 1
			for ; i > 0; i-- {
				k1 = base[i]*m + k1
				if !isPrime3(k1) {
					break
				}
				m *= 10
			}
			if i == 0 {
				arr = append(arr, k)
			}
		}
		d := 2
		i = n - 1
		for d > 0 {
			if i == -1 {
				base = append([]int{d}, base...)
				break
			}
			d += base[i]
			base[i] = d % 10
			d /= 10
			i--
		}
	}
	ans := 0
	fmt.Println(arr)
	for _, v := range arr {
		ans += v
	}
	return ans
}

func isPrime3(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
