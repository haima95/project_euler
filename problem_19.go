package main

import "fmt"

/*   Counting Sundays

You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

*/

func countingSundays() int {
	month := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	weekDay := 1
	year := 1900
	ans := 0
	for year < 2001 {
		for i, day := range month {
			if i == 1 && ((year%4 == 0 && year%100 != 0) || year%400 == 0) {
				day = 29
			}
			fmt.Print(year, "-", i+1, " : ", weekDay, "   ")
			if year > 1900 && weekDay == 0 {
				ans++
			}
			weekDay = (day + weekDay) % 7
		}
		fmt.Println()
		year++
	}
	return ans
}
