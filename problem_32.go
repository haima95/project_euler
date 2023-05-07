package main

import "fmt"

/*		Pandigital products

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example,
the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1
through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.

*/

func test() int {
	ans := 0
	for i := 1; i < 1000000; i++ {
		tmp := make([]bool, 10)
		t := i
		for t > 0 {
			if tmp[t%10] {
				t = -1
				break
			}
			tmp[t%10] = true
			t /= 10
		}
		if t == -1 {
			continue
		}
		for j := i + 1; j < 1000000; j++ {
			t = j
			for t > 0 {
				if tmp[t%10] {
					t = -1
					break
				}
				tmp[t%10] = true
				t /= 10
			}
			if t == -1 {
				continue
			}
			t = j * i
			for t > 0 {
				if tmp[t%10] {
					t = -1
					break
				}
				tmp[t%10] = true
				t /= 10
			}
			if t == -1 {
				continue
			}
			k := 1
			for ; k < 10; k++ {
				if !tmp[k] {
					break
				}
			}
			if k == 10 {
				ans += i * j
				fmt.Println(i, j, ":", j*i)
			}
		}
	}
	return ans
}

func pandigitalProducts() int {
	ans := 0
	kkk := make(map[int]bool)
	for i := 2; i < 10; i++ {
		tmp1 := make([]bool, 10)
		t := i
		for t > 0 {
			if tmp1[t%10] {
				t = -1
				break
			}
			tmp1[t%10] = true
			t /= 10
		}
		if t == -1 {
			continue
		}
		//fmt.Println(i, tmp1)
		for j := 1234; j < 10000; j++ {
			tmp2 := make([]bool, 10)
			for k := 0; k < 10; k++ {
				tmp2[k] = tmp1[k]
			}
			t = j
			for t > 0 {
				if tmp2[t%10] {
					t = -1
					break
				}
				tmp2[t%10] = true
				t /= 10
			}
			if t == -1 {
				continue
			}
			//fmt.Println(i, j, tmp2)
			t = j * i
			for t > 0 {
				if tmp2[t%10] {
					t = -1
					break
				}
				tmp2[t%10] = true
				t /= 10
			}
			if t == -1 || tmp2[0] {
				continue
			}
			k := 1
			for ; k < 10; k++ {
				if !tmp2[k] {
					break
				}
			}
			if k == 10 {
				kkk[i*j] = true
				fmt.Println(i, j, ":", j*i)
			}
		}
	}

	for i := 12; i < 100; i++ {
		tmp1 := make([]bool, 10)
		t := i
		for t > 0 {
			if tmp1[t%10] {
				t = -1
				break
			}
			tmp1[t%10] = true
			t /= 10
		}
		if t == -1 {
			continue
		}
		for j := 123; j < 1000; j++ {
			tmp2 := make([]bool, 10)
			for k := 0; k < 10; k++ {
				tmp2[k] = tmp1[k]
			}
			t = j
			for t > 0 {
				if tmp2[t%10] {
					t = -1
					break
				}
				tmp2[t%10] = true
				t /= 10
			}
			if t == -1 {
				continue
			}
			t = j * i
			for t > 0 {
				if tmp2[t%10] {
					t = -1
					break
				}
				tmp2[t%10] = true
				t /= 10
			}
			if t == -1 || tmp2[0] {
				continue
			}
			k := 1
			for ; k < 10; k++ {
				if !tmp2[k] {
					break
				}
			}
			if k == 10 {
				kkk[i*j] = true
				fmt.Println(i, j, ":", j*i)
			}
		}
	}
	for k, _ := range kkk {
		ans += k
	}
	return ans
}
