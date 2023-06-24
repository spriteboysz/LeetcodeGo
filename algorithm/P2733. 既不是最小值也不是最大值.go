/**
 * Author: Deean
 * Date: 2023-06-24 15:53
 * FileName: algorithm/P2733. 既不是最小值也不是最大值.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func findNonMinOrMax(nums []int) int {
	maximum, minimum := math.MinInt, math.MaxInt
	for _, num := range nums {
		if num > maximum {
			maximum = num
		}
		if num < minimum {
			minimum = num
		}
	}
	for _, num := range nums {
		if num != maximum && num != minimum {
			return num
		}
	}
	return -1
}

func main() {
	fmt.Println(findNonMinOrMax([]int{3, 2, 1, 4}))
}
