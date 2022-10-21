/**
 * Author: Deean
 * Date: 2022-10-21 22:08
 * FileName: sword/剑指 Offer 03. 数组中重复的数字.go
 * Description:
 */

package main

import "fmt"

func findRepeatNumber(nums []int) int {
	hash := map[int]bool{}
	for _, num := range nums {
		if hash[num] == true {
			return num
		} else {
			hash[num] = true
		}
	}
	return -1
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}
