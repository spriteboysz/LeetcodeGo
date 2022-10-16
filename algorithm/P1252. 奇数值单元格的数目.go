/**
 * Author: Deean
 * Date: 2022-10-16 17:35
 * FileName: algorithm/P1252. 奇数值单元格的数目.go
 * Description:
 */

package main

import "fmt"

func oddCells(m int, n int, indices [][]int) int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for _, p := range indices {
		for j := range matrix[p[0]] {
			matrix[p[0]][j]++
		}
		for _, row := range matrix {
			row[p[1]]++
		}
	}
	cnt := 0
	for _, row := range matrix {
		for _, v := range row {
			cnt += v % 2
		}
	}
	return cnt
}

func main() {
	indices := [][]int{{1, 1}, {0, 0}}
	fmt.Println(oddCells(2, 2, indices))
}
