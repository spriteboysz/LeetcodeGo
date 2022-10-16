/**
 * Author: Deean
 * Date: 2022-10-16 16:43
 * FileName: algorithm/P2221. 数组的三角和.go
 * Description:
 */

package main

import "fmt"

func triangularSum(nums []int) int {
	for j := len(nums) - 1; j > 0; j-- {
		for i := range nums[:j] {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
	}
	return nums[0]
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(triangularSum(nums))
}
