/**
 * Author: Deean
 * Date: 2022-10-17 23:12
 * FileName: algorithm/P0062. 不同路径.go
 * Description:
 */

package main

import "fmt"

func uniquePaths(m int, n int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				grid[i][j] = 1
			} else {
				grid[i][j] = grid[i-1][j] + grid[i][j-1]
			}
		}
	}
	return grid[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(7, 3))
}
