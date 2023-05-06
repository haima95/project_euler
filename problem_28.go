package main

/*		Number spiral diagonals

Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:

21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?

*/

func numberSpiralDiagonals(n int) int {
	//tmp := make([][]int, n)
	//for i := 0; i < n; i++ {
	//	tmp[i] = make([]int, n)
	//}
	//rx, ry := n/2, n/2
	//tmp[rx][ry] = 1
	i := 2
	sum := 1
	x, y := 1, 0
	limit := 1
	direct := "down"
	for i < n*n+1 {
		//fmt.Println(direct, x, y, ": ", i)
		//tmp[rx-y][ry+x] = i
		if checkSame(x, y) {
			sum += i
		}
		if direct == "right" {
			x += 1
			if x == limit {
				direct = "down"
			}
		} else if direct == "down" {
			y -= 1
			if -y == limit {
				direct = "left"
			}
		} else if direct == "left" {
			x -= 1
			if -x == limit {
				direct = "up"
			}
		} else {
			y += 1
			if y == limit {
				direct = "right"
				limit += 1
			}
		}
		i++
	}
	//sum2 := 0
	//for i := 0; i < n; i++ {
	//	for j := 0; j < n; j++ {
	//		if i == j || i+j == n-1 {
	//			sum2 += tmp[i][j]
	//		}
	//		//fmt.Printf("%2d ", tmp[i][j])
	//	}
	//	//fmt.Println()
	//}
	//fmt.Println(sum, sum2)
	return sum
}

func checkSame(x, y int) bool {
	return x == y || x == -y
}
