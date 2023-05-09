package main

import "fmt"

/*   Digit factorials

145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: As 1! = 1 and 2! = 2 are not sums they are not included.

*/

// 个人理解，本问题的难度在于如何收敛，即遍历的范围，总不能遍历所有值
// 根据题目要求，可以理解当 9!*n < 10^n 时，10^n之后的值肯定不存在题目要求的值
// 而9!*n是常数项增长，而10^n是指数增长，必然存在n值
// 个人是没具体求n值，而且当 i > i值长度*9!时就退出
func digitFactorials() int {
	ans := 0
	tmp := make([]int, 10)
	tmp[0] = 1
	for i := 1; i < 10; i++ {
		tmp[i] = tmp[i-1] * i
	}
	i := 11
	for {
		sum := 0
		var k []int
		t := i
		for t > 0 {
			sum += tmp[t%10]
			k = append(k, t%10)
			t /= 10
		}
		// 当 i > i值长度*9!时就退出
		if tmp[9]*len(k) < i {
			break
		}
		if sum == i {
			ans += sum
			fmt.Print(i, "== ")
			for j, v := range k {
				if j == len(k)-1 {
					fmt.Println(v, "!")
				} else {
					fmt.Print(v, "! + ")
				}
			}
		}
		i++
	}
	return ans
}
