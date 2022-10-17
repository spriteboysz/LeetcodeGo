/**
 * Author: Deean
 * Date: 2022-10-18 00:00
 * FileName: algorithm/P2057. 值相等的最小索引.go
 * Description:
 */

package main

import "fmt"

func smallestEqual(nums []int) int {
	for i, num := range nums {
		if i%10 == num {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{4, 3, 2, 1}
	fmt.Println(smallestEqual(nums))
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(smallestEqual(nums))
}
