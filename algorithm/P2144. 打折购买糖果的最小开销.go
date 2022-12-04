/**
 * Author: Deean
 * Date: 2022/12/4 16:35
 * FileName: algorithm/P2144. 打折购买糖果的最小开销.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func minimumCost(cost []int) int {
	sort.Ints(cost)
	sum := 0
	for i := len(cost) - 1; i >= 0; i-- {
		if (len(cost)-i)%3 != 0 {
			sum += cost[i]
		}
	}
	return sum
}

func main() {
	fmt.Println(minimumCost([]int{6, 5, 7, 9, 2, 2}))
}
