package main

/*   Prime Pair Sets

The primes 3,7,109 and 673, are quite remarkable. By taking any two primes and concatenating them in any order the
result will always be prime. For example, taking 7 and 109, both 7109 and 1097 are prime. The sum of these four primes,
792, represents the lowest sum for a set of four primes with this property.

Find the lowest sum for a set of five primes for which any two primes concatenate to produce another prime.

*/

func primePairSets() int {
	m := 5
	arr := make([][][]int, 6)
	arr[1] = append(arr[1], []int{2})
	n := 3
	for len(arr[m]) == 0 {
		if isPrime(n) {
			for i := 5; i > 0; i-- {
				for _, v := range arr[i] {
					t := n
					base := 1
					for t > 0 {
						t /= 10
						base *= 10
					}
					if checkPrimeSet(n, base, v) {
						tmp := append([]int{}, v...)
						tmp = append(tmp, n)
						arr[i+1] = append(arr[i+1], tmp)
					}
				}
			}
			arr[1] = append(arr[1], []int{n})
		}
		n += 2
	}
	sum := 0
	for _, v := range arr[m][0] {
		sum += v
	}
	return sum
}

func checkPrimeSet(n, base int, arr []int) bool {
	for _, v := range arr {
		if !isPrime(v*base + n) {
			return false
		}
		t := v
		b := 1
		for t > 0 {
			t /= 10
			b *= 10
		}
		if !isPrime(n*b + v) {
			return false
		}
	}
	return true
}
