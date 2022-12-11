/**
 * Author: Deean
 * Date: 2022/12/11 16:40
 * FileName: algorithm/P0200. 岛屿数量.go
 * Description:
 */

package main

import "fmt"

func numIslands(grid [][]byte) int {
	var dfs func(int, int)
	dfs = func(x, y int) {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		dfs(x-1, y)
		dfs(x, y-1)
		dfs(x+1, y)
		dfs(x, y+1)
	}

	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				cnt++
				dfs(i, j)
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}))
}
