/**
 * Author: Deean
 * Date: 2022/11/15 22:24
 * FileName: algorithm/P1351. 统计有序矩阵中的负数.go
 * Description:
 */

package main

import "fmt"

func countNegatives(grid [][]int) int {
	cnt := 0
	for _, row := range grid {
		for _, val := range row {
			if val < 0 {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(countNegatives([][]int{{1, 2}, {3, 4}}))
}
