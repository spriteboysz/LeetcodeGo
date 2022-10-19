/**
 * Author: Deean
 * Date: 2022-10-19 22:48
 * FileName: algorithm/P1822. 数组元素积的符号.go
 * Description:
 */

package main

import "fmt"

func arraySign(nums []int) int {
	negative := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			negative++
		}
	}
	if negative%2 == 0 {
		return 1
	} else {
		return -1
	}
}

func main() {
	nums := []int{1, 5, 0, 2, -3}
	fmt.Println(arraySign(nums))
	nums = []int{-1, 1, -1, 1, -1}
	fmt.Println(arraySign(nums))
}
