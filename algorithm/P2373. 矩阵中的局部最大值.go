/**
 * Author: Deean
 * Date: 2022-10-13 23:23
 * FileName: algorithm/P2373. 矩阵中的局部最大值.go
 * Description:
 */

package main

import "fmt"

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	for i, row := range grid[:n-2] {
		for j := 0; j < n-2; j++ {
			maximum := 0
			for _, r := range grid[i : i+3] {
				for _, v := range r[j : j+3] {
					if v > maximum {
						maximum = v
					}
				}
			}
			row[j] = maximum
		}
		grid[i] = row[:n-2]
	}
	return grid[:n-2]
}

func main() {
	grid := [][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 2, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1}}
	fmt.Println(largestLocal(grid))
}
