/**
 * Author: Deean
 * Date: 2022/11/19 22:16
 * FileName: algorithm/P2148. 元素计数.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func countElements(nums []int) int {
	maximum, minimum, cnt := math.MinInt32, math.MaxInt32, 0
	for _, num := range nums {
		if maximum < num {
			maximum = num
		}
		if minimum > num {
			minimum = num
		}
	}
	for _, num := range nums {
		if num == minimum || num == maximum {
			cnt++
		}
	}
	return len(nums) - cnt
}

func main() {
	fmt.Println(countElements([]int{-3, 3, 3, 90}))
}
