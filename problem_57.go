package main

// 1 + 1/(1+a/b)
//=1 + 1/((b+a)/b)
//=1 + b/(b+a)
//=(2b+a)/(b+a)

func squareRootConvergents() int {
	i := 2
	numerator := []int{3}
	denominator := []int{2}
	ans := 0
	for i < 1001 {
		tn := sumInt(numerator, denominator)
		td := sumInt(tn, denominator)
		if len(td) > len(tn) {
			//fmt.Println(i, ":", td, tn)
			ans++
		}
		numerator, denominator = td, tn
		i++
	}
	return ans
}

func sumInt(a []int, b []int) []int {
	n, m := len(a), len(b)
	j := 0
	var d []int
	t := 0
	for j < n || j < m {
		if j < n {
			t += a[j]
		}
		if j < m {
			t += b[j]
		}
		d = append(d, t%10)
		t /= 10
		j++
	}
	if t > 0 {
		d = append(d, t)
	}
	return d
}
