package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/*   Coded triangle numbers

The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we
form a word value. For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a triangle
number then we shall call the word a triangle word.

Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common
English words, how many are triangle words?

*/

func codedTriangleNumbers() int {
	ans := 0
	file, _ := os.Open("words.txt")
	content, _ := io.ReadAll(file)
	tmp := strings.Split(string(content), ",")
	for ii, w := range tmp {
		k := 0
		for i := 1; i < len(w)-1; i++ {
			k += int(w[i]-'A') + 1
		}
		if ii < 10 {
			fmt.Println(w, k)
		}
		k *= 2
		for j := 1; j < k; j++ {
			d := j * (j + 1)
			if d > k {
				break
			}
			if d == k {
				ans++
				break
			}
		}
	}
	return ans
}
