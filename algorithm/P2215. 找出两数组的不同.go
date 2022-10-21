/**
 * Author: Deean
 * Date: 2022-10-21 22:29
 * FileName: algorithm/P2215. 找出两数组的不同.go
 * Description:
 */

package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	hash1, hash2 := map[int]bool{}, map[int]bool{}
	for _, num := range nums1 {
		hash1[num] = true
	}
	for _, num := range nums2 {
		hash2[num] = true
	}
	var diff1 []int
	var diff2 []int
	for k := range hash1 {
		if !hash2[k] {
			diff1 = append(diff1, k)
		}
	}
	for k := range hash2 {
		if !hash1[k] {
			diff2 = append(diff2, k)
		}
	}
	return [][]int{diff1, diff2}
}

func main() {
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 1, 2, 2}
	fmt.Println(findDifference(nums1, nums2))
}
