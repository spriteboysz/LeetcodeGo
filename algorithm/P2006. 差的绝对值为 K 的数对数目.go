/**
 * Author: Deean
 * Date: 2022-10-14 00:02
 * FileName: algorithm/P2006. 差的绝对值为 K 的数对数目.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func countKDifference(nums []int, k int) int {
	cnt := 0
	for i, x := range nums {
		for _, y := range nums[:i] {
			if int(math.Abs(float64(x-y))) == k {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	nums := []int{3, 2, 1, 5, 4}
	fmt.Println(countKDifference(nums, 2))
}
