/**
 * Author: Deean
 * Date: 2022/11/20 15:02
 * FileName: algorithm/P2210. 统计数组中峰和谷的数量.go
 * Description:
 */

package main

import "fmt"

func countHillValley(nums []int) int {
	cnt := 0
	for i, n := 0, len(nums); i < n; {
		start, v := i, nums[i]
		for i < n && nums[i] == v {
			i++
		}
		if start > 0 && i < n && nums[start-1] < v == (nums[i] < v) {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(countHillValley([]int{2, 4, 1, 1, 6, 5}))

}
