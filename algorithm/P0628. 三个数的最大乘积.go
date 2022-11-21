/**
 * Author: Deean
 * Date: 2022/11/21 23:34
 * FileName: algorithm/P0628. 三个数的最大乘积.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sort.Ints(nums)
	n := len(nums)
	return max(nums[0]*nums[1]*nums[n-1], nums[n-3]*nums[n-2]*nums[n-1])
}

func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3, 4}))
}
