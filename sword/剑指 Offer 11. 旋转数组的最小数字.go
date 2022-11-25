/**
 * Author: Deean
 * Date: 2022/11/25 22:42
 * FileName: sword/剑指 Offer 11. 旋转数组的最小数字.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func minArray(numbers []int) int {
	minimum := math.MaxInt32
	for _, num := range numbers {
		if num < minimum {
			minimum = num
		}
	}
	return minimum
}

func main() {
	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))
}
