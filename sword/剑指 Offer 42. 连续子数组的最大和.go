/**
 * Author: Deean
 * Date: 2022/12/10 17:13
 * FileName: sword/剑指 Offer 42. 连续子数组的最大和.go
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
