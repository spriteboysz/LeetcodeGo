/**
 * Author: Deean
 * Date: 2022/12/11 16:19
 * FileName: interview/面试题 16.17. 连续数列.go
 * Description:
 */

package main

import "fmt"

func maxSubArray(nums []int) int {
	maximum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > maximum {
			maximum = nums[i]
		}
	}
	return maximum
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
