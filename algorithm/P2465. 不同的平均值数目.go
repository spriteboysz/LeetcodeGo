/**
 * Author: Deean
 * Date: 2022/11/13 23:19
 * FileName: algorithm/P2465. 不同的平均值数目.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	hash := map[int]int{}
	for i := 0; i < n/2; i++ {
		cur := nums[i] + nums[n-i-1]
		hash[cur]++
	}
	return len(hash)
}

func main() {
	fmt.Println(distinctAverages([]int{4, 1, 4, 0, 3, 5}))
}
