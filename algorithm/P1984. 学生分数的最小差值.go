/**
 * Author: Deean
 * Date: 2022/11/16 23:38
 * FileName: algorithm/P1984. 学生分数的最小差值.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	sort.Ints(nums)
	minimum := math.MaxInt32
	for i, num := range nums[:len(nums)-k+1] {
		minimum = min(minimum, nums[i+k-1]-num)
	}
	return minimum
}

func main() {
	fmt.Println(minimumDifference([]int{9, 4, 1, 7}, 2))
}
