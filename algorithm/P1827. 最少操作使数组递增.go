/**
 * Author: Deean
 * Date: 2022-10-16 23:56
 * FileName: algorithm/P1827. 最少操作使数组递增.go
 * Description:
 */

package main

import "fmt"

func minOperations1827(nums []int) int {
	pre := nums[0]
	ret := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > pre {
			pre = nums[i]
		} else {
			ret += pre + 1 - nums[i]
			pre = pre + 1
		}
	}
	return ret
}

func main() {
	nums := []int{1, 5, 2, 4, 1}
	fmt.Println(minOperations1827(nums))
}
