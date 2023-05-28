package main

/*   Powerful digit sum

A googol (10100) is a massive number: one followed by one-hundred zeros; 100100 is almost unimaginably large: one
followed by two-hundred zeros. Despite their size, the sum of the digits in each number is only 1.

Considering natural numbers of the form, ab, where a, b < 100, what is the maximum digital sum?

*/

func powerfulDigitSum() int {
	sum := 1
	for a := 2; a < 100; a++ {
		k := a
		var tmp, base []int
		total := 0
		for k > 0 {
			tmp = append(tmp, k%10)
			base = append(base, k%10)
			k /= 10
			total += k % 10
		}
		if total > sum {
			sum = total
		}
		//fmt.Println(a, "-", 1, ":", tmp)
		for b := 2; b < 100; b++ {
			var tmp1 []int
			d := 0
			for i := 0; i < len(base); i++ {
				d = 0
				for j := 0; j < len(tmp); j++ {
					k = d + base[i]*tmp[j]
					if i+j < len(tmp1) {
						k += tmp1[i+j]
						tmp1[i+j] = k % 10
					} else {
						tmp1 = append(tmp1, k%10)
					}
					d = k / 10
				}
				if d > 0 {
					tmp1 = append(tmp1, d)
				}
			}
			tmp = tmp1
			//fmt.Println(a, "-", b, ":", tmp)
			total = 0
			for _, v := range tmp {
				total += v
			}
			if total > sum {
				sum = total
			}
		}
	}
	return sum
}

// 1 2
// 1 2

// 0 2 4
// 1 2
