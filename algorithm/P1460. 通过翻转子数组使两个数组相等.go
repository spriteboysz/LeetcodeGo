/**
 * Author: Deean
 * Date: 2022-10-17 00:02
 * FileName: algorithm/P1460. 通过翻转子数组使两个数组相等.go
 * Description:
 */

package main

import "fmt"

func canBeEqual(target []int, arr []int) bool {
	var nums [1001]int
	for _, num := range target {
		nums[num]++
	}
	for _, num := range arr {
		nums[num]--
	}
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}
func main() {
	target := []int{1, 2, 3, 4}
	arr := []int{2, 4, 1, 3}
	fmt.Println(canBeEqual(target, arr))
}
