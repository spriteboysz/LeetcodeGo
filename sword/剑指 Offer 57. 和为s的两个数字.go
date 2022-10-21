/**
 * Author: Deean
 * Date: 2022-10-21 22:04
 * FileName: sword/剑指 Offer 57. 和为s的两个数字.go
 * Description:
 */

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] < target {
			left++
		} else if nums[left]+nums[right] > target {
			right--
		} else {
			return []int{nums[left], nums[right]}
		}
	}
	return []int{}
}

func main() {
	nums := []int{10, 26, 30, 31, 47, 60}
	fmt.Println(twoSum(nums, 40))
}
