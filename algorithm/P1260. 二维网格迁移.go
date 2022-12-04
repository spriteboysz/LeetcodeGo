/**
 * Author: Deean
 * Date: 2022/12/4 17:34
 * FileName: algorithm/P1260. 二维网格迁移.go
 * Description:
 */

package main

import "fmt"

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	shifted := make([][]int, m)
	for i := range shifted {
		shifted[i] = make([]int, n)
	}
	for i, row := range grid {
		for j, num := range row {
			index := (i*n + k + j) % (m * n)
			shifted[index/n][index%n] = num
		}
	}
	return shifted
}

func main() {
	fmt.Println(shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
}
