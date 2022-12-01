/**
 * Author: Deean
 * Date: 2022/12/1 23:07
 * FileName: algorithm/P2475. 数组中不等三元组的数目.go
 * Description:
 */

package main

import "fmt"

func unequalTriplets(nums []int) int {
	cnt, n := 0, len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				continue
			}
			for k := j + 1; k < n; k++ {
				if nums[i] != nums[k] && nums[j] != nums[k] {
					cnt++
				}
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(unequalTriplets([]int{4, 4, 2, 4, 3}))
}
