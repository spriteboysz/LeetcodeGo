/**
 * Author: Deean
 * Date: 2022/12/14 22:07
 * FileName: algorithm/P0912. 排序数组.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func sortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func main() {
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}
