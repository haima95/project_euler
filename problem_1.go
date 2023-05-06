package main

/*
Problem 1: Multiples of 3 or 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

func findMultiples(n int) int {
	tmp := make(map[int]bool)
	base := 3
	i := 1
	m := base * i
	for m < n {
		tmp[m] = true
		i++
		m = base * i
	}
	base = 5
	i = 1
	m = base * i
	for m < n {
		tmp[m] = true
		i++
		m = base * i
	}
	ans := 0
	for k, _ := range tmp {
		ans += k
	}
	return ans
}
