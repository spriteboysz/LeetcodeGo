/**
 * Author: Deean
 * Date: 2022-10-13 23:12
 * FileName: algorithm/P1470. 重新排列数组.go
 * Description:
 */

package main

import "fmt"

func shuffle(nums []int, n int) []int {
	var nums2 []int
	for i := 0; i < n; i++ {
		nums2 = append(nums2, nums[i])
		nums2 = append(nums2, nums[n+i])
	}
	return nums2
}

func main() {
	nums := []int{1, 2, 3, 4, 4, 3, 2, 1}
	fmt.Println(shuffle(nums, 4))
}
