/**
 * Author: Deean
 * Date: 2022/11/25 22:45
 * FileName: algorithm/P0154. 寻找旋转排序数组中的最小值 II.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func findMin(nums []int) int {
	minimum := math.MaxInt32
	for _, num := range nums {
		if num < minimum {
			minimum = num
		}
	}
	return minimum
}

func main() {
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
}
