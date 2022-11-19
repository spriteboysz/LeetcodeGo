/**
 * Author: Deean
 * Date: 2022/11/19 16:28
 * FileName: algorithm/P1608. 特殊数组的特征值.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func specialArray(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, n := 1, len(nums); i <= n; i++ {
		if nums[i-1] >= i && (i == n || nums[i] < i) {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(specialArray([]int{0, 4, 3, 0, 4}))
}
