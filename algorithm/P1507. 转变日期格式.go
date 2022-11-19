/**
 * Author: Deean
 * Date: 2022/11/19 20:03
 * FileName: algorithm/P1507. 转变日期格式.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reformatDate(date string) string {
	months := map[string]int{
		"Jan": 1,
		"Feb": 2,
		"Mar": 3,
		"Apr": 4,
		"May": 5,
		"Jun": 6,
		"Jul": 7,
		"Aug": 8,
		"Sep": 9,
		"Oct": 10,
		"Nov": 11,
		"Dec": 12,
	}
	dates := strings.Split(date, " ")
	day, _ := strconv.Atoi(dates[0][:len(dates[0])-2])
	month := months[dates[1]]
	year := dates[2]
	return fmt.Sprintf("%s-%02d-%02d", year, month, day)
}

func main() {
	fmt.Println(reformatDate("6th Jun 1933"))
}
