package main

import "fmt"

/*		Reciprocal cycles

A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10
are given:

1/2	= 	0.5
1/3	= 	0.(3)
1/4	= 	0.25
1/5	= 	0.2
1/6	= 	0.1(6)
1/7	= 	0.(142857)
1/8	= 	0.125
1/9	= 	0.(1)
1/10	= 	0.1
Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.

*/

// 可以优化，求出素数，因为素数肯定会有循环
func reciprocalCycles() int {
	max, ans := 6, 7
	for d := 11; d < 1000; d += 2 {
		tmp := make(map[string]bool)
		s, k := 0, 1
		for k != 0 {
			k *= 10
			s, k = k/d, k%d
			key := fmt.Sprintf("%d_%d", s, k)
			if _, ok := tmp[key]; ok {
				break
			}
			tmp[key] = true
		}
		//fmt.Println(d, ": ", tmp)
		if len(tmp) > max {
			max = len(tmp)
			ans = d
		}
	}
	return ans
}
