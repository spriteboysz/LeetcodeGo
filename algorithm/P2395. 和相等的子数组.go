/**
 * Author: Deean
 * Date: 2022/12/4 16:05
 * FileName: algorithm/P2395. 和相等的子数组.go
 * Description:
 */

package main

import "fmt"

func findSubarrays(nums []int) bool {
	hash := map[int]bool{}
	for i := 1; i < len(nums); i++ {
		s := nums[i-1] + nums[i]
		if hash[s] {
			return true
		}
		hash[s] = true
	}
	return false
}

func main() {
	fmt.Println(findSubarrays([]int{4, 2, 4}))
}
