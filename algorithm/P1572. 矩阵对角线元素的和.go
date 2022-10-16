/**
 * Author: Deean
 * Date: 2022-10-16 16:23
 * FileName: algorithm/P1572. 矩阵对角线元素的和.go
 * Description:
 */

package main

import "fmt"

func diagonalSum(mat [][]int) int {
	sum, n := 0, len(mat)
	for i := range mat {
		sum += mat[i][i] + mat[i][n-1-i]
	}
	if n%2 == 1 {
		sum -= mat[n/2][n/2]
	}
	return sum
}

func main() {
	mat := [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}
	fmt.Println(diagonalSum(mat))
	mat = [][]int{{5}}
	fmt.Println(diagonalSum(mat))
}
