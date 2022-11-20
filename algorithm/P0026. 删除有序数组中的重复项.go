/**
 * Author: Deean
 * Date: 2022/11/20 21:38
 * FileName: algorithm/P0026. 删除有序数组中的重复项.go
 * Description:
 */

package main

import "fmt"

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {
	fmt.Println(removeDuplicates2([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
