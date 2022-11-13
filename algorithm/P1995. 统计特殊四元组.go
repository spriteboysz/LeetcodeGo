/**
 * Author: Deean
 * Date: 2022/11/13 22:56
 * FileName: algorithm/P1995. 统计特殊四元组.go
 * Description:
 */

package main

import "fmt"

func countQuadruplets(nums []int) int {
	cnt := 0
	for a := range nums {
		for b := a + 1; b < len(nums); b++ {
			for c := b + 1; c < len(nums); c++ {
				for d := c + 1; d < len(nums); d++ {
					if nums[a]+nums[b]+nums[c] == nums[d] {
						cnt++
					}
				}
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(countQuadruplets([]int{1, 1, 1, 3, 5}))
}
