package main

/*   Number letter counts

If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19
letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?


NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one
hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.

*/

func numberLetterCounts() int {
	down_10 := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	down_20 := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen",
		"eighteen", "nineteen"}
	base_10 := []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	base_100 := []string{"hundred"}
	tmp10, tmp100, ans := 0, 0, len("onethousand")
	for i := 0; i < 9; i++ {
		tmp10 += len(down_10[i])
	}
	tmp100 += tmp10
	for i := 0; i < 10; i++ {
		tmp100 += len(down_20[i])
	}
	for i := 0; i < 8; i++ {
		tmp100 += len(base_10[i])*10 + tmp10
	}
	for i := 0; i < 9; i++ {
		ans += 100*(len(down_10[i])+len(base_100[0])) + 99*3 + tmp100
	}
	return ans + tmp100
}
