/**
 * Author: Deean
 * Date: 2022/12/8 23:37
 * FileName: sword/剑指 Offer 49. 丑数.go
 * Description:
 */

package main

import "fmt"

func nthUglyNumber(n int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	a, b, c := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		n1, n2, n3 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(min(n1, n2), n3)
		if dp[i] == n1 {
			a++
		}
		if dp[i] == n2 {
			b++
		}
		if dp[i] == n3 {
			c++
		}
	}
	return dp[n-1]
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
