/**
 * Author: Deean
 * Date: 2022/12/4 16:56
 * FileName: sword/剑指 Offer II 012. 左右两边子数组的和相等.go
 * Description:
 */

package main

import "fmt"

func pivotIndex(nums []int) int {
	sum, pre := 0, 0
	for _, num := range nums {
		sum += num
	}
	for i, num := range nums {
		if 2*pre+num == sum {
			return i
		}
		pre += num
	}
	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
