package main

import "fmt"

func combinatoricSelections() int {
	max := 1000000
	ans := 0
	for n := 23; n <= 100; n++ {
		a := 1
		// 因为组合是左右对称，使用求一半就能得出结果，因为中间是最高点，然后左右下降，并对称
		// 即  C(n,m) == C(n,n-m)
		for j := 1; j <= n/2; j++ {
			a = a * (n - j + 1) / j
			if n == 23 && j == 10 {
				fmt.Println(a)
			}
			if a > max {
				if n%2 == 0 {
					ans += 2*(n/2-j) + 1
				} else {
					ans += 2 * (n/2 - j + 1)
				}
				break
			}
		}
	}
	return ans
}

// 0 1 2 3 4
// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23
