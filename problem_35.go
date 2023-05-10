package main

import "fmt"

/*		Circular primes

The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?

*/

func circularPrimes(m int) int {
	result := make(map[int]bool)
	visited := make(map[int]bool)
	result[2] = true
	for i := 3; i < m; i += 2 {
		if visited[i] {
			continue
		}
		match := true
		t := i
		d := 1
		n := 0
		for t > 0 {
			if t%10 == 0 {
				match = false
				break
			}
			d *= 10
			t /= 10
			n++
		}
		if !match {
			continue
		}
		d /= 10
		var arr []int
		j := 0
		t = i
		for j < n {
			if !isPrime2(t) {
				match = false
			}
			arr = append(arr, t)
			visited[t] = true
			t = t%10*d + t/10
			j++
		}
		if match {
			for _, v := range arr {
				result[v] = true
			}
		}
	}
	fmt.Println(result)
	return len(result)
}

func isPrime2(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
