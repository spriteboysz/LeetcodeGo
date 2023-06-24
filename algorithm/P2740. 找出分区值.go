/**
 * Author: Deean
 * Date: 2023-06-24 15:25
 * FileName: algorithm/P2740. 找出分区值.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
	"sort"
)

func findValueOfPartition(nums []int) int {
	sort.Ints(nums)
	minimum := math.MaxInt
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] < minimum {
			minimum = nums[i] - nums[i-1]
		}
	}
	return minimum
}

func main() {
	fmt.Println(findValueOfPartition([]int{1, 3, 2, 4}))
}
