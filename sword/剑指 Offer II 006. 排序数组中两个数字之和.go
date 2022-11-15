/**
 * Author: Deean
 * Date: 2022/11/15 23:11
 * FileName: sword/剑指 Offer II 006. 排序数组中两个数字之和.go
 * Description:
 */

package main

import "fmt"

func twoSum2(numbers []int, target int) []int {
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
	return []int{left, right}
}

func main() {
	fmt.Println(twoSum2([]int{2, 3, 5}, 7))
}
