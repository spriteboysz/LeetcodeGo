/**
 * Author: Deean
 * Date: 2022/12/7 23:23
 * FileName: algorithm/P0215. 数组中的第K个最大元素.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
