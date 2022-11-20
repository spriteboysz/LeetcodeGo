/**
 * Author: Deean
 * Date: 2022/11/20 17:26
 * FileName: algorithm/P0976. 三角形的最大周长.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	for i := n - 1; i >= 2; i-- {
		if nums[i] < nums[i-1]+nums[i-2] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))
}
