/**
 * Author: Deean
 * Date: 2022-10-11 22:55
 * FileName: algorithm/P0598. 范围求和 II.go
 * Description:
 */

package main

import "fmt"

func maxCount(m int, n int, ops [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	minRow, minCol := m, n
	for _, op := range ops {
		minRow = min(minRow, op[0])
		minCol = min(minCol, op[1])
	}
	return minRow * minCol
}

func main() {
	ops := [][]int{{2, 2}, {3, 3}}
	fmt.Println(maxCount(3, 3, ops))
}
