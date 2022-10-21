/**
 * Author: Deean
 * Date: 2022-10-21 22:21
 * FileName: algorithm/P0867. 转置矩阵.go
 * Description:
 */

package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	matrix2 := make([][]int, n)
	for i := range matrix2 {
		matrix2[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix2[j][i] = matrix[i][j]
		}
	}
	return matrix2
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(transpose(matrix))
}
