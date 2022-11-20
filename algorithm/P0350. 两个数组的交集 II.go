/**
 * Author: Deean
 * Date: 2022/11/20 17:59
 * FileName: algorithm/P0350. 两个数组的交集 II.go
 * Description:
 */

package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	cnt1, cnt2 := map[int]int{}, map[int]int{}
	for _, num := range nums1 {
		cnt1[num]++
	}
	for _, num := range nums2 {
		cnt2[num]++
	}
	target := []int{}
	for k, v := range cnt1 {
		for i := 0; i < min(v, cnt2[k]); i++ {
			target = append(target, k)
		}
	}
	return target
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}
