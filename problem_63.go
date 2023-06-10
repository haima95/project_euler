package main

import "fmt"

/*   Powerful Digit Counts

The 5-digit number 16807 = 7^5 is also a fifth power. Similarly, the 9-digit number 134217728 = 8^9 is a ninth power.

How many n-digit positive integers exist which are also an n th power?

*/
// 9^n == 10^(n-1)
func powerfulDigitCounts() int {
	/* 根据公式，可知道，当 n 满足9^n < 10^(n-1)后，大于等于n的值都不会满足条件，求出n就知道上限值是几，避免无限计算
	   因为python支持大数，所以先用python求出n

	>>> n = 1
	>>> k = 9
	>>> t = 10
	>>> while True:
	...     print(f'{k*10} - {t} : {k*10 > t}')
	...     if k*10 < t:
	...             break
	...     k *= 9
	...     t *= 10
	...     n += 1
	...

	得到 n = 22

	*/
	tmp := make([][]int, 9)
	for i := 0; i < 9; i++ {
		tmp[i] = []int{1}
	}
	ans := 0
	n := 1
	for n < 22 {
		for i := 0; i < 9; i++ {
			d := 0
			for j := 0; j < len(tmp[i]); j++ {
				d = tmp[i][j]*(9-i) + d
				tmp[i][j] = d % 10
				d /= 10
			}
			if d > 0 {
				tmp[i] = append(tmp[i], d)
			}
			if (n == 5 && 9-i == 7) || (n == 9 && 9-i == 8) {
				fmt.Println(tmp[i], ": ", len(tmp[i]) == n)
			}
			if len(tmp[i]) == n {
				ans++
			}
		}
		n++
	}
	return ans
}
