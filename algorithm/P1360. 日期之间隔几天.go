/**
 * Author: Deean
 * Date: 2022/11/30 23:36
 * FileName: algorithm/P1360. 日期之间隔几天.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func daysBetweenDates(date1 string, date2 string) int {
	months := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	isLeap := func(year int) int {
		if (year%100 != 0 && year%4 == 0) || year%400 == 0 {
			return 1
		}
		return 0
	}
	calc := func(date string) int {
		d := strings.Split(date, "-")
		year, _ := strconv.Atoi(d[0])
		month, _ := strconv.Atoi(d[1])
		day, _ := strconv.Atoi(d[2])

		days := day
		for y := 1971; y < year; y++ {
			days += isLeap(y) + 365
		}
		for m := 1; m < month; m++ {
			if m == 2 {
				days += isLeap(year) + 28
			} else {
				days += months[m]
			}
		}
		return days
	}
	return abs(calc(date1) - calc(date2))
}

func main() {
	fmt.Println(daysBetweenDates("2020-01-15", "2019-12-31"))
}
