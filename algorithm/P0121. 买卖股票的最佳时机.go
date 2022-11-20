/**
 * Author: Deean
 * Date: 2022/11/20 17:13
 * FileName: algorithm/P0121. 买卖股票的最佳时机.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	maximum, minimum := math.MinInt32, math.MaxInt32
	for _, price := range prices {
		minimum = min(minimum, price)
		maximum = max(maximum, price-minimum)
	}
	return maximum
}

func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
