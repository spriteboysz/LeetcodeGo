/**
 * Author: Deean
 * Date: 2022-10-21 21:38
 * FileName: algorithm/P0977. 有序数组的平方.go
 * Description:
 */

package main

import "fmt"

func sortedSquares(nums []int) []int {
	sorted := make([]int, len(nums))
	for left, right, i := 0, len(nums)-1, len(nums)-1; left <= right; i-- {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			sorted[i] = nums[left] * nums[left]
			left++
		} else {
			sorted[i] = nums[right] * nums[right]
			right--
		}
	}
	return sorted
}

func main() {
	nums := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(nums))
}
