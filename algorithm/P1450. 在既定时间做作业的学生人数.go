/**
 * Author: Deean
 * Date: 2022-10-15 22:08
 * FileName: algorithm/P1450. 在既定时间做作业的学生人数.go
 * Description:
 */

package main

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	cnt := 0
	for i := 0; i < len(startTime); i++ {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			cnt++
		}
	}
	return cnt
}

func main() {
	start := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	end := []int{10, 10, 10, 10, 10, 10, 10, 10, 10}
	fmt.Println(busyStudent(start, end, 5))
}
