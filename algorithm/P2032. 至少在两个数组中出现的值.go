/**
 * Author: Deean
 * Date: 2022-10-21 22:35
 * FileName: algorithm/P2032. 至少在两个数组中出现的值.go
 * Description:
 */

package main

import "fmt"

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	hash := map[int]int{}
	for _, num := range nums1 {
		hash[num] |= 1
	}
	for _, num := range nums2 {
		hash[num] |= 2
	}
	for _, num := range nums3 {
		hash[num] |= 4
	}
	var target []int
	for k, v := range hash {
		if v == 3 || v == 5 || v == 6 || v == 7 {
			target = append(target, k)
		}
	}
	return target
}

func main() {
	nums1 := []int{3, 1}
	nums2 := []int{3, 2}
	nums3 := []int{1, 2}
	fmt.Println(twoOutOfThree(nums1, nums2, nums3))
}
