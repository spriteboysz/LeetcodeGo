/**
 * Author: Deean
 * Date: 2023-10-15 22:54
 * FileName: LCP/LCR 076. 数组中的第 K 个最大元素.go
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
