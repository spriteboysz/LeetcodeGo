/**
 * Author: Deean
 * Date: 2022/11/20 21:43
 * FileName: algorithm/P2239. 找到最接近 0 的数字.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func findClosestNumber(nums []int) int {
	abs := func(x int) int {
		if x >= 0 {
			return x
		}
		return -x
	}
	sort.Slice(nums, func(i, j int) bool {
		ai, aj := abs(nums[i]), abs(nums[j])
		if ai == aj {
			return nums[i] > nums[j]
		}
		return ai < aj
	})
	return nums[0]
}

func main() {
	fmt.Println(findClosestNumber([]int{2, -1, 1}))
}
