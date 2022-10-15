/**
 * Author: Deean
 * Date: 2022-10-15 21:58
 * FileName: algorithm/P1389. 按既定顺序创建目标数组.go
 * Description:
 */

package main

import "fmt"

func createTargetArray(nums []int, index []int) []int {
	var target = make([]int, len(nums))
	for k, i := range index {
		copy(target[i+1:], target[i:])
		target[i] = nums[k]
	}
	return target
}

func main() {
	nums := []int{1, 2, 3, 4, 0}
	index := []int{0, 1, 2, 3, 0}
	fmt.Println(createTargetArray(nums, index))
}
