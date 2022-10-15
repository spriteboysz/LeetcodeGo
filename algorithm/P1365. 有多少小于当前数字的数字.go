/**
 * Author: Deean
 * Date: 2022-10-15 22:19
 * FileName: algorithm/P1365. 有多少小于当前数字的数字.go
 * Description:
 */

package main

import (
	"fmt"
)

func smallerNumbersThanCurrent(nums []int) []int {
	smaller := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[i] {
				smaller[i]++
			}
		}
	}
	return smaller
}

func main() {
	nums := []int{6, 5, 4, 8}
	fmt.Println(smallerNumbersThanCurrent(nums))
}
