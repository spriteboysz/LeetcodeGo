/**
 * Author: Deean
 * Date: 2022/11/20 22:23
 * FileName: algorithm/P2016. 增量元素之间的最大差值.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func maximumDifference(nums []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maximum, minimum := math.MinInt32, math.MaxInt32
	for _, num := range nums {
		minimum = min(minimum, num)
		maximum = max(maximum, num-minimum)
	}
	if maximum == 0 {
		return -1
	}
	return maximum
}

func main() {
	fmt.Println(maximumDifference([]int{1, 5, 2, 10}))
	fmt.Println(maximumDifference([]int{9, 4, 3, 2}))
}
