package main

import "fmt"

/*   Integer right triangles

If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions
for p = 120.

{20,48,52}, {24,45,51}, {30,40,50}

For which value of p ≤ 1000, is the number of solutions maximised?

*/

func integerRightTriangles() int {
	max := 0
	index := 0
	for p := 3; p < 1001; p++ {
		n := 0
		tmp := [][]int{}
		for a := 1; a < p; a++ {
			for b := a; b < p; b++ {
				c := p - a - b
				if c > b && c*c == a*a+b*b {
					n++
					tmp = append(tmp, []int{a, b, c})
				}
			}
		}
		if p == 120 {
			fmt.Println(tmp)
		}
		if n > max {
			max = n
			index = p
		}
	}
	return index
}
