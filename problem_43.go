package main

import "fmt"

/*   Sub-string divisibility

The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the digits 0 to 9 in some order,
but it also has a rather interesting sub-string divisibility property.

Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:

d2d3d4=406 is divisible by 2
d3d4d5=063 is divisible by 3
d4d5d6=635 is divisible by 5
d5d6d7=357 is divisible by 7
d6d7d8=572 is divisible by 11
d7d8d9=728 is divisible by 13
d8d9d10=289 is divisible by 17
Find the sum of all 0 to 9 pandigital numbers with this property.

*/

func subStringDivisibility() int {
	divid := map[int]int{
		4:  2,
		5:  3,
		6:  5,
		7:  7,
		8:  11,
		9:  13,
		10: 17,
	}
	ans := 0
	tmp := make([]bool, 10)
	for i := 1; i < 10; i++ {
		tmp[i] = true
		createNumber(i, 1, tmp, divid, &ans)
		tmp[i] = false
	}
	return ans
}

func createNumber(base int, index int, tmp []bool, divid map[int]int, ans *int) {
	if index > 3 {
		d := base % 1000
		if d%divid[index] != 0 {
			return
		}
		if index == 10 {
			fmt.Println(base)
			*ans += base
		}
	}
	for i := 0; i < 10; i++ {
		if tmp[i] {
			continue
		}
		tmp[i] = true
		createNumber(base*10+i, index+1, tmp, divid, ans)
		tmp[i] = false
	}
}
