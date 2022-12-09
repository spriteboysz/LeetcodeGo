/**
 * Author: Deean
 * Date: 2022/12/8 23:24
 * FileName: interview/面试题 10.11. 峰与谷.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) {
	sort.Ints(nums)
	for i := 1; i < len(nums); i += 2 {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}
	fmt.Println(nums)
}

func main() {
	wiggleSort([]int{5, 3, 1, 2, 3})
}
