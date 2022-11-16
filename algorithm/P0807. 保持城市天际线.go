/**
 * Author: Deean
 * Date: 2022-10-11 23:06
 * FileName: algorithm/P0807. 保持城市天际线.go
 * Description:
 */

package main

import "fmt"

func maxIncreaseKeepingSkyline(grid [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n := len(grid)
	rowMax, colMax := make([]int, n), make([]int, n)
	for i, row := range grid {
		for j, val := range row {
			rowMax[i] = max(rowMax[i], val)
			colMax[j] = max(colMax[j], val)
		}
	}
	cnt := 0
	for i, row := range grid {
		for j, val := range row {
			cnt += min(rowMax[i], colMax[j]) - val
		}
	}
	return cnt
}

func main() {
	grid := [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
}
