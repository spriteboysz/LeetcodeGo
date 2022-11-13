/**
 * Author: Deean
 * Date: 2022/11/13 17:23
 * FileName: algorithm/P1991. 找到数组的中间位置.go
 * Description:
 */

package main

import "fmt"

func findMiddleIndex(nums []int) int {
	total, prefix := 0, 0
	for _, num := range nums {
		total += num
	}
	for i, num := range nums {
		if prefix*2+num == total {
			return i
		}
		prefix += num
	}
	return -1
}

func main() {
	fmt.Println(findMiddleIndex([]int{2, 3, -1, 8, 4}))
}
