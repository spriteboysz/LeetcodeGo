/**
 * Author: Deean
 * Date: 2022/12/11 17:45
 * FileName: sword/剑指 Offer II 104. 排列的数目.go
 * Description:
 */

package main

import "fmt"

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3, 4}, 7))
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}
