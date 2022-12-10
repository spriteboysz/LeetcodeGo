/**
 * Author: Deean
 * Date: 2022/12/10 17:17
 * FileName: algorithm/P0540. 有序数组中的单一元素.go
 * Description:
 */

package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
}
