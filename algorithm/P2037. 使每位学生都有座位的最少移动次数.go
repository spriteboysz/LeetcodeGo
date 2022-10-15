/**
 * Author: Deean
 * Date: 2022-10-15 23:15
 * FileName: algorithm/P2037. 使每位学生都有座位的最少移动次数.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	abs := func(num int) int {
		if num >= 0 {
			return num
		}
		return -num
	}
	sort.Ints(seats)
	sort.Ints(students)
	cnt := 0
	for i, seat := range seats {
		cnt += abs(students[i] - seat)
	}
	return cnt
}

func main() {
	seats := []int{4, 1, 5, 9}
	students := []int{1, 3, 2, 6}
	fmt.Println(minMovesToSeat(seats, students))
}
