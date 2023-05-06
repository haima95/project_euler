package main

/*	Digit fifth powers

Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

1634 = 14 + 64 + 34 + 44
8208 = 84 + 24 + 04 + 84
9474 = 94 + 44 + 74 + 44
As 1 = 14 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.

*/

func digitFifthPowers(n int) int {
	total := n + 1
	for i := 0; i < n; i++ {
		total *= 10
	}
	ans := 0
	for i := 2; i < total; i++ {
		sum := 0
		t := i
		for t > 0 {
			base := t % 10
			t /= 10
			d := 1
			for j := 0; j < n; j++ {
				d *= base
			}
			sum += d
		}
		if sum == i {
			//fmt.Println(i)
			ans += sum
		}
	}
	return ans
}
