/**
 * Author: Deean
 * Date: 2022-10-18 23:53
 * FileName: algorithm/P0349. 两个数组的交集.go
 * Description:
 */

package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	hash := map[int]int{}
	for _, num := range nums1 {
		hash[num] |= 1
	}
	for _, num := range nums2 {
		hash[num] |= 2
	}
	var nums []int
	for k, v := range hash {
		if v == 3 {
			nums = append(nums, k)
		}
	}
	return nums
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}
