/**
 * Author: Deean
 * Date: 2022/11/12 21:21
 * FileName: algorithm/P2293. 极大极小游戏.go
 * Description:
 */

package main

import "fmt"

func minMaxGame(nums []int) int {
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
	for len(nums) > 1 {
		cur := make([]int, len(nums)/2)
		for i := range cur {
			if i%2 == 0 {
				cur[i] = min(nums[i*2], nums[i*2+1])
			} else {
				cur[i] = max(nums[i*2], nums[i*2+1])
			}
		}
		nums = cur
	}
	return nums[0]
}

func main() {
	fmt.Println(minMaxGame([]int{1, 3, 5, 2, 4, 8, 2, 2}))
}
