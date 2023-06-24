/**
 * Author: Deean
 * Date: 2023-06-24 15:43
 * FileName: algorithm/P2735. 收集巧克力.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func minCost(nums []int, x int) int64 {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(nums)
	sum := make([]int, len(nums))
	for i := range sum {
		sum[i] = i * x
	}
	for i, mn := range nums {
		for j := i; j < n+i; j++ {
			mn = min(mn, nums[j%n])
			sum[j-i] += mn
		}
	}
	minimum := math.MaxInt
	for _, s := range sum {
		minimum = min(minimum, s)
	}
	return int64(minimum)
}

func main() {
	fmt.Println(minCost([]int{20, 1, 15}, 5))
}
