package main

import "fmt"

/*  Amicable numbers

Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The
proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.

*/

func amicableNumbers() int {
	visited := make(map[int]bool)
	sum := 0
	for i := 2; i < 10000; i++ {
		if _, ok := visited[i]; ok {
			continue
		}
		sum1 := check(i)
		visited[i] = true
		if _, ok := visited[sum1]; ok {
			continue
		}
		sum2 := check(sum1)
		if sum2 == i {
			fmt.Println(sum1, sum2)
			sum += sum2 + sum1
			visited[sum1] = true
		}
	}
	return sum
}

func check(n int) int {
	sum := 1
	j := 2
	for ; j*j < n; j++ {
		if n%j == 0 {
			sum += j + n/j
		}
	}
	if j*j == n {
		sum += j
	}
	return sum
}
