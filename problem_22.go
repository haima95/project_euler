package main

import (
	"io"
	"os"
	"sort"
	"strings"
)

/*     Names scores

Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names,
begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value
by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the
938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

*/

func namesScores() int {
	sum := 0
	file, _ := os.Open("name.txt")
	data, _ := io.ReadAll(file)
	data2 := string(data)
	tmp := strings.Split(data2, ",")
	//fmt.Println(len(tmp))
	sort.Strings(tmp)
	for i, v := range tmp {
		k := 0
		for _, c := range []byte(v[1 : len(v)-1]) { // 去掉头尾两个双引号
			k += int(c-'A') + 1
		}
		//fmt.Println(i+1, v[1:len(v)-1], ": ", k)
		sum += (i + 1) * k
	}
	return sum
}
