/**
 * Author: Deean
 * Date: 2022/11/14 22:53
 * FileName: sword/P1636. 按照频率将数组升序排序.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func frequencySort2(nums []int) []int {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num] += 1
	}
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		hx, hy := hash[x], hash[y]
		if hx == hy {
			return x > y
		} else {
			return hx < hy
		}
	})
	return nums
}

func main() {
	fmt.Println(frequencySort2([]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}))
}
