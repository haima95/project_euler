package main

/*   Lexicographic permutations

A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3
and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The
lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

*/

func lexicographicPermutations(num int) string {
	num -= 1
	tmp := make([]int, 10)
	tmp[1] = 1
	for i := 2; i < 10; i++ {
		tmp[i] = tmp[i-1] * i
	}
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var ans []int
	for i := 9; i > 0; i-- {
		j := 0
		for ; j < len(arr) && num >= tmp[i]; j++ {
			num -= tmp[i]
		}
		ans = append(ans, arr[j])
		list := arr[:j]
		list = append(list, arr[j+1:]...)
		arr = list
	}
	ans = append(ans, arr...)
	ss := make([]byte, len(ans))
	for i := 0; i < len(ans); i++ {
		ss[i] = '0' + byte(ans[i])
	}
	return string(ss)
}
