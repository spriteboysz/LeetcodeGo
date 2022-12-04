/**
 * Author: Deean
 * Date: 2022/12/4 11:39
 * FileName: algorithm/P1863. 找出所有子集的异或总和再求和.go
 * Description:
 */

package main

import "fmt"

func subsetXORSum(nums []int) int {
	xor := 0
	var dfs func(s, now int)
	dfs = func(s, now int) {
		xor += now
		if s == len(nums) {
			return
		}
		for i := s; i < len(nums); i++ {
			dfs(i+1, now^nums[i])
		}
	}
	dfs(0, 0)
	return xor
}

func main() {
	fmt.Println(subsetXORSum([]int{5, 1, 6}))
}
