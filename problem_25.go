package main

/*	1000-digit Fibonacci number

The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

*/

func get1000digitFibonacciNumber(n int) int {
	index := 2
	f1 := []int{1}
	f2 := []int{1}
	for {
		var tmp []int
		i, d := 0, 0
		for ; i < len(f1); i++ {
			d += f1[i] + f2[i]
			tmp = append(tmp, d%10)
			d /= 10
		}
		for ; i < len(f2); i++ {
			d += f2[i]
			tmp = append(tmp, d%10)
			d /= 10
		}
		if d > 0 {
			tmp = append(tmp, d)
		}
		f1, f2 = f2, tmp
		index++
		if len(f2) >= n {
			return index
		}
	}
	return -1
}
