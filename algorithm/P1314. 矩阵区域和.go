/**
 * Author: Deean
 * Date: 2022-11-07 23:30
 * FileName: algorithm/P1314. 矩阵区域和.go
 * Description:
 */

package main

import "fmt"

func matrixBlockSum(mat [][]int, k int) [][]int {
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	m, n := len(mat), len(mat[0])
	block := make([][]int, 0)
	for i, row := range mat {
		cur := make([]int, 0)
		for j := range row {
			val := 0
			for di := max(0, i-k); di <= min(i+k, m-1); di++ {
				for dj := max(0, j-k); dj <= min(j+k, n-1); dj++ {
					val += mat[di][dj]
				}
			}
			cur = append(cur, val)
		}
		block = append(block, cur)
	}
	return block
}

func main() {
	fmt.Println(matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
}
