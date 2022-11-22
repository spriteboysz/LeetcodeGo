/**
 * Author: Deean
 * Date: 2022/11/21 23:38
 * FileName: algorithm/P1646. 获取生成数组中的最大值.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func getMaximumGenerated(n int) int {
	nums := make([]int, n+1)
	nums[1] = 1
	for i := 1; i <= n/2; i++ {
		nums[2*i] = nums[i]
		if 2*i+1 <= n {
			nums[2*i+1] = nums[i] + nums[i+1]
		}
	}
	sort.Ints(nums)
	return nums[n]
}

func main() {
	fmt.Println(getMaximumGenerated(7))
}
