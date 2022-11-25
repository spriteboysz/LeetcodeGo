/**
 * Author: Deean
 * Date: 2022/11/25 22:31
 * FileName: algorithm/P0034. 在排序数组中查找元素的第一个和最后一个位置.go
 * Description:
 */

package main

import "fmt"

func searchRange(nums []int, target int) []int {
	left, right := -1, -1
	for i, num := range nums {
		if left == -1 && num == target {
			left = i
			right = i
		}
		if right != -1 && num == target {
			right = i
		}
		if right != -1 && num != target {
			break
		}
	}
	return []int{left, right}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}
