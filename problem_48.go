package main

/*   Self powers

The series, 11 + 22 + 33 + ... + 1010 = 10405071317.

Find the last ten digits of the series, 11 + 22 + 33 + ... + 10001000.

*/

func selfPowers(n int) int {
	sum := make([]int, 10)
	for i := 1; i <= n; i++ {
		var base []int
		m := 0
		k := i
		for k > 0 && m < 10 {
			base = append(base, k%10)
			k /= 10
			m++
		}
		tmp := make([]int, 10)
		d := 0
		for j := 0; j < i; j++ {
			if j == 0 {
				for j1 := 0; j1 < m; j1++ {
					tmp[j1] = base[j1]
				}
			} else {
				tmp2 := make([]int, 10)
				for j1 := 0; j1 < m; j1++ {
					d = 0
					for j2 := 0; j2 < 10; j2++ {
						if j1 == 0 {
							d += tmp[j2] * base[j1]
							tmp2[j2] = d % 10
							d /= 10
						} else {
							if j2+j1 < 10 {
								d += tmp[j2]*base[j1] + tmp2[j2+j1]
								tmp2[j2+j1] = d % 10
								d /= 10
							} else {
								break
							}
						}
					}
				}
				tmp = tmp2
			}
			//fmt.Println(i, "-", j, ": ", tmp)

		}
		//fmt.Println(i, ": ", tmp)
		d = 0
		for j := 0; j < 10; j++ {
			d += sum[j] + tmp[j]
			sum[j] = d % 10
			d /= 10
		}
		//fmt.Println(i, ": ", tmp)
	}
	ans := 0
	b := 1
	for i := 0; i < 10; i++ {
		ans = ans + sum[i]*b
		b *= 10
	}
	return ans
}
