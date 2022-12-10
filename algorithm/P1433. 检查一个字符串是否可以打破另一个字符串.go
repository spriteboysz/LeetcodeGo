/**
 * Author: Deean
 * Date: 2022/12/9 23:18
 * FileName: algorithm/P1433. 检查一个字符串是否可以打破另一个字符串.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func checkIfCanBreak(s1 string, s2 string) bool {
	nums1, nums2 := []byte(s1), []byte(s2)
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})

	flag1, flag2 := true, true
	for i := 0; i < len(nums1); i++ {
		if nums1[i] < nums2[i] {
			flag1 = false
			break
		}
	}
	for i := 0; i < len(nums2); i++ {
		if nums2[i] < nums1[i] {
			flag2 = false
			break
		}
	}
	return flag1 || flag2
}

func main() {
	fmt.Println(checkIfCanBreak("leetcodee", "interview"))
}
