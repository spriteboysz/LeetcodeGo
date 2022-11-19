/**
 * Author: Deean
 * Date: 2022/11/19 19:53
 * FileName: algorithm/P1893. 检查是否区域内所有整数都被覆盖.go
 * Description:
 */

package main

import "fmt"

func isCovered(ranges [][]int, left int, right int) bool {
	nums := make([]int, 51)
	for _, r := range ranges {
		start, end := r[0], r[1]
		for i := start; i <= end; i++ {
			nums[i]++
		}
	}
	for i := left; i <= right; i++ {
		if nums[i] == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isCovered([][]int{{1, 10}, {10, 20}}, 21, 21))
}
