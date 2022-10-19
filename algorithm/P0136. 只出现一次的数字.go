/**
 * Author: Deean
 * Date: 2022-10-19 23:18
 * FileName: algorithm/P0136. 只出现一次的数字.go
 * Description:
 */

package main

import "fmt"

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
