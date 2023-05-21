package main

/*   Distinct primes factors

The first two consecutive numbers to have two distinct prime factors are:

14 = 2 × 7
15 = 3 × 5

The first three consecutive numbers to have three distinct prime factors are:

644 = 2² × 7 × 23
645 = 3 × 5 × 43
646 = 2 × 17 × 19.

Find the first four consecutive integers to have four distinct prime factors each. What is the first of these numbers?

*/

func distinctPrimesFactors(n int) int {
	//doPrime(14, 0, 2, 2, map[int]bool{2: true, 3: true, 5: true, 7: true})
	//return 0
	d := 4
	tmp := make([]int, n)
	primeArr := make(map[int]bool)
	primeArr[2] = true
	primeArr[3] = true
	i := 0
	//fmt.Println(d, tmp)
	for i < n {
		isprime := true
		for j := 2; j*j <= d; j++ {
			if d%j == 0 {
				isprime = false
				break
			}
		}
		if doPrime(d, 0, n, 2, primeArr) {
			if i == 0 || tmp[i-1] == d-1 {
				tmp[i] = d
				i++
			} else {
				tmp[0] = d
				i = 1
			}
		}
		if isprime {
			primeArr[d] = true
		}
		d++
		//fmt.Println(d, tmp)
	}
	return tmp[0]
}

func doPrime(d, m, n, s int, primeArr map[int]bool) bool {
	//fmt.Println(d, m, n, s, primeArr)
	if d == 1 {
		if m == n {
			return true
		}
		return false
	}
	for i := s; i <= d; i++ {
		if d%i == 0 && primeArr[i] {
			k := d
			for k%i == 0 {
				k /= i
			}
			if doPrime(k, m+1, n, i+1, primeArr) {
				return true
			}
		}
	}
	return false
}
