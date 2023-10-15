/**
 * Author: Deean
 * Date: 2023-10-14 21:44
 * FileName: LCP/LCR 006. 两数之和 II - 输入有序数组.go
 * Description:
 */

package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] < target {
			left += 1
		} else if numbers[left]+numbers[right] > target {
			right -= 1
		} else {
			return []int{left, right}
		}
	}
	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 4, 6, 10}, 8))
}
