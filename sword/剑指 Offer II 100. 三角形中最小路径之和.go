/**
 * Author: Deean
 * Date: 2022/12/7 22:39
 * FileName: sword/剑指 Offer II 100. 三角形中最小路径之和.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func minimumTotal(triangle [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	minimum := math.MaxInt32
	for i := 0; i < n; i++ {
		minimum = min(minimum, dp[n-1][i])
	}
	return minimum
}

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}
