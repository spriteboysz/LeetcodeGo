/**
 * Author: Deean
 * Date: 2022-10-10 23:54
 * FileName: algorithm/P1920. 基于排列构建数组.go
 * Description:
 */

package main

import "fmt"

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i, num := range nums {
		ans[i] = nums[num]
	}
	return ans
}

func main() {
	nums := []int{5, 0, 1, 2, 3, 4}
	fmt.Println(buildArray(nums))
}
