/**
 * Author: Deean
 * Date: 2022/11/13 16:21
 * FileName: algorithm/P0746. 使用最小花费爬楼梯.go
 * Description:
 */

package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i < n+1; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
