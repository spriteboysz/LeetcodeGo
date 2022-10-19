/**
 * Author: Deean
 * Date: 2022-10-19 22:21
 * FileName: algorithm/P2319. 判断矩阵是否是一个 X 矩阵.go
 * Description:
 */

package main

import "fmt"

func checkXMatrix(grid [][]int) bool {
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i == n-1-j {
				if grid[i][j] == 0 {
					return false
				}
			} else {
				if grid[i][j] != 0 {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	grid := [][]int{{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2}}
	fmt.Println(checkXMatrix(grid))
}
