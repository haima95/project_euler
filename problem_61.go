package main

import "fmt"

func cyclicalFigurateNumbers() int {
	triangle := make(map[int][]int)
	square := make(map[int][]int)
	pentagonal := make(map[int][]int)
	hexagonal := make(map[int][]int)
	heptagonal := make(map[int][]int)
	octagonal := make(map[int][]int)
	i := 1
	no := 0
	for no != 6 {
		no = 0
		t := i * (i + 1) / 2
		if t > 1000 && t < 10000 {
			triangle[t/100] = append(triangle[t/100], t%100)
		} else if t > 9999 {
			no++
		}
		t = i * i
		if t > 1000 && t < 10000 {
			square[t/100] = append(square[t/100], t%100)
		} else if t > 9999 {
			no++
		}
		t = i * (3*i - 1) / 2
		if t > 1000 && t < 10000 {
			pentagonal[t/100] = append(pentagonal[t/100], t%100)
		} else if t > 9999 {
			no++
		}
		t = i * (2*i - 1)
		if t > 1000 && t < 10000 {
			hexagonal[t/100] = append(hexagonal[t/100], t%100)
		} else if t > 9999 {
			no++
		}
		t = i * (5*i - 3) / 2
		if t > 1000 && t < 10000 {
			heptagonal[t/100] = append(heptagonal[t/100], t%100)
		} else if t > 9999 {
			no++
		}
		t = i * (3*i - 2)
		if t > 1000 && t < 10000 {
			octagonal[t/100] = append(octagonal[t/100], t%100)
		} else if t > 9999 {
			no++
		}
		//fmt.Println(i, ":", no)
		i++
	}
	fmt.Println("count")
	for _, v := range triangle[81] {
		if v == 28 {
			fmt.Println("1")
			break
		}
	}
	for _, v := range square[28] {
		if v == 82 {
			fmt.Println("2")
			break
		}
	}
	for _, v := range pentagonal[82] {
		if v == 81 {
			fmt.Println("3")
			break
		}
	}
	//triangle = map[int][]int{81: []int{28}}
	//square = map[int][]int{28: []int{82}}
	//pentagonal = map[int][]int{82: []int{81}}
	tmp22 := []map[int][]int{triangle, square, pentagonal, hexagonal, heptagonal, octagonal}
	var tmp1 [][]map[int][]int
	for i, v := range tmp22 {
		kk := append([]map[int][]int{}, v)
		tmp3 := append([]map[int][]int{}, tmp22[:i]...)
		tmp3 = append(tmp3, tmp22[i+1:]...)
		group(kk, tmp3, &tmp1)

	}
	fmt.Println("kk: ", len(tmp1))
	for _, tmp := range tmp1 {
		m := 6
		ans := make([]int, m)
		for k, arr := range tmp[0] {
			for _, v := range arr {
				ans[0] = k*100 + v
				if doCheck(tmp, 1, v, ans, m) {
					sum := 0
					for i := 0; i < m; i++ {
						sum += ans[i]
					}
					fmt.Println(ans, ":", sum)
				}
			}
		}
	}
	return -1
}

func doCheck(tmp []map[int][]int, index int, pre int, ans []int, m int) bool {
	if index == m {
		return ans[index-1]%100 == ans[0]/100
	}
	for _, v := range tmp[index][pre] {
		ans[index] = pre*100 + v
		if doCheck(tmp, index+1, v, ans, m) {
			return true
		}
	}
	return false
}

func group(tmp, arr []map[int][]int, tmp1 *[][]map[int][]int) {
	if len(arr) == 0 {
		*tmp1 = append(*tmp1, tmp)
	}
	for i, v := range arr {
		kk := append([]map[int][]int{}, tmp...)
		kk = append(kk, v)
		tmp3 := append([]map[int][]int{}, arr[:i]...)
		tmp3 = append(tmp3, arr[i+1:]...)
		group(kk, tmp3, tmp1)
	}
}
