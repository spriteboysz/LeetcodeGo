/**
 * Author: Deean
 * Date: 2022/11/13 23:55
 * FileName: algorithm/P1217. 玩筹码.go
 * Description:
 */

package main

import "fmt"

func minCostToMoveChips(position []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	cnt := [2]int{}
	for _, p := range position {
		cnt[p%2]++
	}
	return min(cnt[0], cnt[1])
}

func main() {
	fmt.Println(minCostToMoveChips([]int{2, 2, 2, 3, 3}))
}
