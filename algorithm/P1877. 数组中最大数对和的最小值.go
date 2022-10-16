/**
 * Author: Deean
 * Date: 2022-10-16 13:24
 * FileName: algorithm/P1877. 数组中最大数对和的最小值.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	minimum := nums[0] + nums[len(nums)-1]
	for i := 0; i < len(nums)/2; i++ {
		cur := nums[i] + nums[len(nums)-i-1]
		if cur > minimum {
			minimum = cur
		}
	}
	return minimum
}

func main() {
	nums := []int{3, 5, 4, 2, 4, 6}
	fmt.Println(minPairSum(nums))
}
