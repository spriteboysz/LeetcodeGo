/**
 * Author: Deean
 * Date: 2022-11-08 22:47
 * FileName: algorithm/P0260. 只出现一次的数字 III.go
 * Description:
 */

package main

import "fmt"

func singleNumberIII(nums []int) []int {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num] += 1
	}
	single := []int{}
	for num, occ := range hash {
		if occ == 1 {
			single = append(single, num)
		}
	}
	return single
}

func main() {
	fmt.Println(singleNumberIII([]int{1, 2, 1, 3, 2, 5}))
}
