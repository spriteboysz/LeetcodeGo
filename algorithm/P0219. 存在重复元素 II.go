/**
 * Author: Deean
 * Date: 2022/11/25 23:28
 * FileName: algorithm/P0219. 存在重复元素 II.go
 * Description:
 */

package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	for i, num := range nums {
		for j := i + 1; j-i <= k && j < len(nums); j++ {
			if num == nums[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}
