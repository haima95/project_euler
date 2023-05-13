package main

import "fmt"

/*	Pandigital multiples

Take the number 192 and multiply it by each of 1, 2, and 3:

192 × 1 = 192
192 × 2 = 384
192 × 3 = 576
By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product
of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645,
which is the concatenated product of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer
with (1,2, ... , n) where n > 1?

*/

func pandigitalMultiples() int {
	max := 0
	for i := 2; i < 10; i++ {
		k := 9 / i
		base := 1
		for k > 1 {
			base *= 10
			k--
		}
		for j := base; j < base*10; j++ {
			tmp := make([]int, 10)
			var arr []int
			tmp[0] = 1
			ans := 0
			k = 1
			for ; k <= i; k++ {
				t := j * k
				for t > 0 {
					if tmp[t%10] == 1 {
						break
					}
					tmp[t%10] = 1
					t /= 10
					ans *= 10
				}
				if t > 0 {
					break
				}
				arr = append(arr, j*k)
				ans += j * k
			}
			if k <= i {
				continue
			}
			k = 1
			for ; k < 10; k++ {
				if tmp[k] == 0 {
					break
				}
			}
			if k == 10 {
				fmt.Println(j, i, arr, ": ", ans)
				if ans > max {
					max = ans
				}
			}
		}
	}
	return max
}
