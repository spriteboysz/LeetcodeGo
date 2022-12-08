/**
 * Author: Deean
 * Date: 2022/12/7 23:21
 * FileName: sword/剑指 Offer II 076. 数组中的第 k 大的数字.go
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
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
