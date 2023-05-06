package main

import "fmt"

/* 	Factorial digit sum

n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!

*/

func factorialDigitSum(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	arr := []int{1}
	for i := 2; i <= n; i++ {
		tmp := make([]int, len(arr))
		for j, v := range arr {
			d := v * i
			if d == 0 {
				continue
			}
			for k := j; k < len(tmp); k++ {
				d += tmp[k]
				tmp[k] = d % 10
				d /= 10
			}
			for d > 0 {
				tmp = append(tmp, d%10)
				d /= 10
			}
		}
		arr = tmp
		fmt.Println(i, ": ", tmp)
	}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
