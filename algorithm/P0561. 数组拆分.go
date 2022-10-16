/**
 * Author: Deean
 * Date: 2022-10-16 21:10
 * FileName: algorithm/P0561. 数组拆分.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func main() {
	nums := []int{6, 2, 6, 5, 1, 2}
	fmt.Println(arrayPairSum(nums))
}
