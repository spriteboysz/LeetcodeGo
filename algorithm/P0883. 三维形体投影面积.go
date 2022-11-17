/**
 * Author: Deean
 * Date: 2022/11/17 23:03
 * FileName: algorithm/P0883. 三维形体投影面积.go
 * Description:
 */

package main

import (
	"fmt"
)

func projectionArea(grid [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	area := 0
	for i, row := range grid {
		maxRow, maxCol := 0, 0
		for j, val := range row {
			if val != 0 {
				area++
			}
			maxRow = max(maxRow, val)
			maxCol = max(maxCol, grid[j][i])
		}
		area += maxRow + maxCol
	}
	return area
}

func main() {
	fmt.Println(projectionArea([][]int{{1, 2}, {3, 4}}))
}
