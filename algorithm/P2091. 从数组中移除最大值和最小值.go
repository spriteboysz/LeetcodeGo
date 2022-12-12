/**
 * Author: Deean
 * Date: 2022/12/12 23:03
 * FileName: algorithm/P2091. 从数组中移除最大值和最小值.go
 * Description:
 */

package main

import "fmt"

func minimumDeletions(nums []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	i, j := 0, 0
	for k, num := range nums {
		if num < nums[i] {
			i = k
		}
		if num > nums[j] {
			j = k
		}
	}
	if i > j {
		i, j = j, i
	}
	return min(min(j+1, len(nums)-i), i+1+len(nums)-j)
}

func main() {
	fmt.Println(minimumDeletions([]int{2, 10, 7, 5, 4, 1, 8, 6}))
}
