/**
 * Author: Deean
 * Date: 2023-10-16 22:49
 * FileName: LCP/LCR 135. 报数.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func countNumbers(cnt int) []int {
	nums := []int{}
	for i, n := 1, math.Pow10(cnt); i < (int)(n); i++ {
		nums = append(nums, i)
	}
	return nums
}

func main() {
	fmt.Println(countNumbers(2))
}
