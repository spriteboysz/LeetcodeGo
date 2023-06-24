/**
 * Author: Deean
 * Date: 2023-06-24 17:19
 * FileName: algorithm/P2706. 购买两块巧克力.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	if money >= prices[0]+prices[1] {
		return money - prices[0] - prices[1]
	}
	return money
}

func main() {
	fmt.Println(buyChoco([]int{1, 2, 2}, 3))
}
