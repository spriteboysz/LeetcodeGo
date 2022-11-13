/**
 * Author: Deean
 * Date: 2022/11/13 17:46
 * FileName: algorithm/P0283. 移动零.go
 * Description:
 */

package main

import "fmt"

func moveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
	fmt.Println(nums)
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}
