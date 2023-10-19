/**
 * Author: Deean
 * Date: 2023-10-16 23:36
 * FileName: algorithm/P0054. 螺旋矩阵.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func spiralOrder(matrix [][]int) []int {
	m, nums := len(matrix), []int{}
	if m == 0 {
		return nums
	}
	n := len(matrix[0])
	if n == 0 {
		return nums
	}
	left, right, top, bottom := 0, n-1, 0, m-1
	for {
		for i := left; i <= right; i++ {
			nums = append(nums, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			nums = append(nums, matrix[i][right])
		}
		right--
		if left > right {
			break
		}
		for i := right; i >= left; i-- {
			nums = append(nums, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}
		for i := bottom; i >= top; i-- {
			nums = append(nums, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return nums
}

func main() {
	fmt.Println(spiralOrder(lib.Str2Array2D("[[1,2,3],[8,9,4],[7,6,5]]")))
}
