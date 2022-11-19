/**
 * Author: Deean
 * Date: 2022/11/19 19:48
 * FileName: algorithm/P0027. 移除元素.go
 * Description:
 */

package main

import "fmt"

func removeElement(nums []int, val int) int {
	left := 0
	for _, num := range nums {
		if num != val {
			nums[left] = num
			left++
		}
	}
	return left
}

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
