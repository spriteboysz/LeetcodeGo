/**
 * Author: Deean
 * Date: 2022/11/26 22:32
 * FileName: algorithm/P0704. 二分查找.go
 * Description:
 */

package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		num := nums[mid]
		if num > target {
			right = mid - 1
		} else if num < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 5))
}
