/**
 * Author: Deean
 * Date: 2022/11/17 23:35
 * FileName: algorithm/P1154. 一年中的第几天.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
)

func dayOfYear(date string) int {
	MONTH := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	days, _ := strconv.Atoi(date[8:])
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		MONTH[1]++
	}
	for _, m := range MONTH[:month-1] {
		days += MONTH[m]
	}
	return days
}

func main() {
	fmt.Println(dayOfYear("2019-01-09"))
}
