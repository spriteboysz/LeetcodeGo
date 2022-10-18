/**
 * Author: Deean
 * Date: 2022-10-18 23:18
 * FileName: algorithm/P0442. 数组中重复的数据.go
 * Description:
 */

package main

import "fmt"

func findDuplicates(nums []int) []int {
	var duplicates []int
	set := map[int]bool{}
	for _, num := range nums {
		if set[num] {
			duplicates = append(duplicates, num)
		} else {
			set[num] = true
		}
	}
	return duplicates
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(findDuplicates(nums))
}
