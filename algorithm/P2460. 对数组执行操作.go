/**
 * Author: Deean
 * Date: 2022/11/12 21:36
 * FileName: algorithm/P2460. 对数组执行操作.go
 * Description:
 */

package main

import "fmt"

func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	return nums
}

func main() {
	fmt.Println(applyOperations([]int{1, 2, 2, 1, 1, 0}))
}
