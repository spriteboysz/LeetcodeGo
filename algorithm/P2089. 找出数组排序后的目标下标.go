/**
 * Author: Deean
 * Date: 2022-10-16 17:46
 * FileName: algorithm/P2089. 找出数组排序后的目标下标.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	var indices []int
	for i, num := range nums {
		if num == target {
			indices = append(indices, i)
		}
	}
	return indices
}

func main() {
	nums := []int{1, 2, 5, 2, 3}
	fmt.Println(targetIndices(nums, 3))
}
