/**
 * Author: Deean
 * Date: 2022/12/4 16:18
 * FileName: algorithm/P1800. 最大升序子数组和.go
 * Description:
 */

package main

import "fmt"

func maxAscendingSum(nums []int) int {
	maximum := 0
	for i, n := 0, len(nums); i < n; {
		sum := nums[i]
		for i++; i < n && nums[i] > nums[i-1]; i++ {
			sum += nums[i]
		}
		if sum > maximum {
			maximum = sum
		}
	}
	return maximum
}

func main() {
	fmt.Println(maxAscendingSum([]int{12, 17, 15, 13, 10, 11, 12}))
}
