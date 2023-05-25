package main

/*   Permuted multiples

It can be seen that the number, 125874, and its double, 251748, contain exactly the same digits, but in a different order.

Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.

*/

func permutedMultiples() int {
	i := 11
	for {
		tmp := make(map[int]int)
		k := i
		for k > 0 {
			tmp[k%10]++
			k /= 10
		}
		nomatch := true
		for j := 2; j < 7 && nomatch; j++ {
			k = j * i
			tmp1 := make(map[int]int)
			for k > 0 {
				tmp1[k%10]++
				k /= 10
			}
			if len(tmp1) != len(tmp) {
				nomatch = false
				break
			}

			for k1, v1 := range tmp {
				if tmp1[k1] != v1 {
					nomatch = false
					break
				}
			}
		}
		if nomatch {
			return i
		}
		i++
	}
	return -1
}
