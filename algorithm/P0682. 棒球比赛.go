/**
 * Author: Deean
 * Date: 2022-11-09 21:57
 * FileName: algorithm/P0682. 棒球比赛.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	points := []int{}
	ans := 0
	for _, op := range operations {
		n := len(points)
		switch op[0] {
		case '+':
			ans += points[n-1] + points[n-2]
			points = append(points, points[n-1]+points[n-2])
		case 'D':
			ans += points[n-1] * 2
			points = append(points, 2*points[n-1])
		case 'C':
			ans -= points[n-1]
			points = points[:len(points)-1]
		default:
			pt, _ := strconv.Atoi(op)
			ans += pt
			points = append(points, pt)
		}
	}
	return ans
}

func main() {
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
}
