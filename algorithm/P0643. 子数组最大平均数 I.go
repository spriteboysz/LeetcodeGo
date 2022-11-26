/**
 * Author: Deean
 * Date: 2022/11/26 22:14
 * FileName: algorithm/P0643. 子数组最大平均数 I.go
 * Description:
 */

package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sum := 0
	for _, num := range nums[:k] {
		sum += num
	}
	maximum := sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		maximum = max(maximum, sum)
	}
	return float64(maximum) / float64(k)
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}
