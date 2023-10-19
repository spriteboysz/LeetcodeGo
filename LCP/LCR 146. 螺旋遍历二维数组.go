/**
 * Author: Deean
 * Date: 2023-10-16 23:27
 * FileName: LCP/LCR 146. 螺旋遍历二维数组.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func spiralArray(array [][]int) []int {
	m, nums := len(array), []int{}
	if m == 0 {
		return nums
	}
	n := len(array[0])
	if n == 0 {
		return nums
	}
	left, right, top, bottom := 0, n-1, 0, m-1
	for {
		for i := left; i <= right; i++ {
			nums = append(nums, array[top][i])
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			nums = append(nums, array[i][right])
		}
		right--
		if left > right {
			break
		}
		for i := right; i >= left; i-- {
			nums = append(nums, array[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}
		for i := bottom; i >= top; i-- {
			nums = append(nums, array[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return nums
}

func main() {
	fmt.Println(spiralArray(lib.Str2Array2D("[[1,2,3],[8,9,4],[7,6,5]]")))
}
