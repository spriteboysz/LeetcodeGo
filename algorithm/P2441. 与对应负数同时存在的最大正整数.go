/**
 * Author: Deean
 * Date: 2022-11-10 22:45
 * FileName: algorithm/P2441. 与对应负数同时存在的最大正整数.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func findMaxK(nums []int) int {
	hash := map[int]bool{}
	for _, num := range nums {
		hash[num] = true
	}
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if hash[-nums[i]] {
			return nums[i]
		}
	}
	return -1
}

func main() {
	fmt.Println(findMaxK([]int{-1, 10, 6, 7, -7, 1}))
}
