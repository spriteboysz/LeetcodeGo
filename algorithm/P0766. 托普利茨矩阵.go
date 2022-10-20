/**
 * Author: Deean
 * Date: 2022-10-20 23:28
 * FileName: algorithm/P0766. 托普利茨矩阵.go
 * Description:
 */

package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	fmt.Println(isToeplitzMatrix(matrix))
}
