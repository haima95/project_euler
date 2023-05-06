package main

/*		Coin sums

In the United Kingdom the currency is made up of pound (£) and pence (p). There are eight coins in general circulation:

1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), and £2 (200p).
It is possible to make £2 in the following way:

1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?

*/

func coinSums(n int) int {
	base := []int{200, 100, 50, 20, 10, 5, 2, 1}
	sum := 0
	for i, v := range base {
		k := v
		tmp := n
		for tmp >= k {
			tmp -= k
			d := countCoins(tmp, base[i+1:])
			if d > 0 {
				sum += d
			}
		}
	}
	return sum
}

func countCoins(n int, base []int) int {
	//fmt.Println(n, base)
	if n == 0 {
		return 1
	}
	if len(base) == 0 {
		return -1
	}
	if len(base) == 1 {
		if n%base[0] == 0 {
			return 1
		} else {
			return -1
		}
	}
	sum := 0
	for i, v := range base {
		k := v
		tmp := n
		for tmp >= k {
			tmp -= k
			d := countCoins(tmp, base[i+1:])
			if d > 0 {
				sum += d
			}
		}
	}
	return sum
}
