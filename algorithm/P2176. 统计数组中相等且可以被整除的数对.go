/**
 * Author: Deean
 * Date: 2022-10-16 17:26
 * FileName: algorithm/P2176. 统计数组中相等且可以被整除的数对.go
 * Description:
 */

package main

import "fmt"

func countPairs(nums []int, k int) int {
	cnt := 0
	for i, a := range nums {
		for j := i + 1; j < len(nums); j++ {
			if a == nums[j] && (i*j)%k == 0 {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	nums := []int{3, 1, 2, 2, 2, 1, 3}
	fmt.Println(countPairs(nums, 2))
}
