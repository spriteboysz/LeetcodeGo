/**
 * Author: Deean
 * Date: 2022/11/15 23:15
 * FileName: algorithm/P0167. 两数之和 II - 输入有序数组.go
 * Description:
 */

package main

import "fmt"

func twoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] < target {
			left++
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			break
		}
	}
	return []int{left + 1, right + 1}
}

func main() {
	fmt.Println(twoSumII([]int{2, 7, 11, 15}, 13))
}
