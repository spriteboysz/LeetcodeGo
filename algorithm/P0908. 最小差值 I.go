/**
 * Author: Deean
 * Date: 2022/11/17 23:13
 * FileName: algorithm/P0908. 最小差值 I.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func smallestRangeI(nums []int, k int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	minimum, maximum := math.MaxInt32, -1
	for _, num := range nums {
		if num > maximum {
			maximum = num
		}
		if num < minimum {
			minimum = num
		}
	}
	return max(0, maximum-k-minimum-k)
}

func main() {
	fmt.Println(smallestRangeI([]int{1, 3, 6}, 3))
}
