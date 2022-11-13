/**
 * Author: Deean
 * Date: 2022/11/13 22:47
 * FileName: algorithm/P2022. 将一维数组转变成二维数组.go
 * Description:
 */

package main

import "fmt"

func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return [][]int{}
	}
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	k := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = original[k]
			k++
		}
	}
	return matrix
}

func main() {
	fmt.Println(construct2DArray([]int{1, 2, 3}, 1, 3))
}
