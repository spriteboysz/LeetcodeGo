/**
 * Author: Deean
 * Date: 2022/11/13 16:05
 * FileName: algorithm/P0566. 重塑矩阵.go
 * Description:
 */

package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	mat2 := make([][]int, r)
	for i := range mat2 {
		mat2[i] = make([]int, c)
	}
	for i := 0; i < n*m; i++ {
		mat2[i/c][i%c] = mat[i/n][i%n]
	}
	return mat2
}

func main() {
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4))
}
