/**
 * Author: Deean
 * Date: 2022/11/16 22:27
 * FileName: algorithm/P0485. 最大连续 1 的个数.go
 * Description:
 */

package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	cnt, maxCnt := 0, 0
	for _, num := range nums {
		if num == 1 {
			cnt++
		} else {
			maxCnt = max(maxCnt, cnt)
			cnt = 0
		}
	}
	return max(maxCnt, cnt)
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
