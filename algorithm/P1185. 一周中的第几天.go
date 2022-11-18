/**
 * Author: Deean
 * Date: 2022/11/18 23:15
 * FileName: algorithm/P1185. 一周中的第几天.go
 * Description:
 */

package main

import "fmt"

func dayOfTheWeek(day int, month int, year int) string {
	var week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}
	days := day
	days += 365*(year-1971) + (year-1969)/4
	for _, d := range monthDays[:month-1] {
		days += d
	}
	if month >= 3 && (year%400 == 0 || year%4 == 0 && year%100 != 0) {
		days++
	}
	return week[(days+3)%7]
}

func main() {
	fmt.Println(dayOfTheWeek(18, 7, 1999))
}
