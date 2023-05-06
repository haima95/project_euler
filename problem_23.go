package main

import "fmt"

/*    Non-abundant sums

A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the
sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum
exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two
abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as
the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is
known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.

*/

func non_abundantSums() int {
	max := 28123
	abundantNums := []int{}
	for i := 1; i <= max; i++ {
		tmp := 1
		j := 2
		for ; j*j < i; j++ {
			if i%j == 0 {
				tmp += j + i/j
			}
		}
		if j*j == i {
			tmp += j
		}
		if tmp > i {
			abundantNums = append(abundantNums, i)
		}
	}
	fmt.Println(abundantNums[:10])
	arr := make(map[int]bool)
	for i := 0; i < len(abundantNums); i++ {
		for j := i; j < len(abundantNums); j++ {
			arr[abundantNums[i]+abundantNums[j]] = true
		}
	}
	fmt.Println(arr[24])
	sum := 0
	for i := 1; i <= max; i++ {
		if arr[i] {
			continue
		}
		sum += i
	}
	return sum
}
