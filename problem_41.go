package main

import "fmt"

/*   Pandigital prime

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once. For example,
2143 is a 4-digit pandigital and is also prime.

What is the largest n-digit pandigital prime that exists?

*/

func pandigitalPrime() int {
	max := 0
	arr := []int{1}
	for i := 2; i < 5; i++ {
		var tmp []int
		for j, v := range arr {
			d := 1
			for d <= v {
				d1, d2 := v/d, v%d
				k := d1*10*d + i*d + d2
				d *= 10
				if k%2 == 1 && isPrime(k) {
					if k > max {
						max = k
					}
				}
				tmp = append(tmp, k)
			}
			k := i*d + arr[j]
			if k%2 == 1 && isPrime(k) {
				if k > max {
					max = k
				}
			}
			tmp = append(tmp, k)
		}
		arr = tmp
		fmt.Println(arr)
	}
	return max
}
