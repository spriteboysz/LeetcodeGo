/**
 * Author: Deean
 * Date: 2022/11/28 22:18
 * FileName: algorithm/P0035. 搜索插入位置.go
 * Description:
 */

package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right, index := 0, len(nums)-1, len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if target <= nums[mid] {
			index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return index
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
