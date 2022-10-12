/**
 * Author: Deean
 * Date: 2022-10-12 23:35
 * FileName: algorithm/P1929. 数组串联.go
 * Description:
 */

package main

import "fmt"

func getConcatenation(nums []int) []int {
	for _, num := range nums {
		nums = append(nums, num)
	}
	return nums
}

func main() {
	nums := []int{1, 3, 2, 1}
	fmt.Println(getConcatenation(nums))
}
