/**
 * Author: Deean
 * Date: 2023-06-24 17:03
 * FileName: algorithm/P2717. 半有序排列.go
 * Description:
 */

package main

import "fmt"

func semiOrderedPermutation(nums []int) int {
	index1, index2, n := -1, -1, len(nums)
	for i, num := range nums {
		if num == 1 {
			index1 = i
		}
		if num == n {
			index2 = i
		}
	}
	diff := index1 + n - 1 - index2
	if index1 < index2 {
		return diff
	}
	return diff - 1
}

func main() {
	fmt.Println(semiOrderedPermutation([]int{2, 1, 4, 3}))
}
