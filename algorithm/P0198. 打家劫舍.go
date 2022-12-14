/**
 * Author: Deean
 * Date: 2022/12/14 22:30
 * FileName: algorithm/P0198. 打家劫舍.go
 * Description:
 */

package main

import "fmt"

func rob(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
