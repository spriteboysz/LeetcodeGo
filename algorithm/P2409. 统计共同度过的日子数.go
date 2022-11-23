/**
 * Author: Deean
 * Date: 2022/11/22 23:17
 * FileName: algorithm/P2409. 统计共同度过的日子数.go
 * Description:
 */

package main

import "fmt"

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	min := func(a, b string) string {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b string) string {
		if a > b {
			return a
		}
		return b
	}
	var months = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	calc := func(s string) (days int) {
		for _, day := range months[:s[0]&15*10+s[1]&15-1] {
			days += day
		}
		return days + int(s[3]&15*10+s[4]&15)
	}
	days := calc(min(leaveAlice, leaveBob)) - calc(max(arriveAlice, arriveBob)) + 1
	if days < 0 {
		return 0
	}
	return days
}

func main() {
	fmt.Println(countDaysTogether("08-15", "08-18", "08-16", "08-19"))
}
