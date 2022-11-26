/**
 * Author: Deean
 * Date: 2022/11/25 23:25
 * FileName: algorithm/P0217. 存在重复元素.go
 * Description:
 */

package main

import "fmt"

func containsDuplicate(nums []int) bool {
	hash := map[int]bool{}
	for _, num := range nums {
		if hash[num] {
			return true
		}
		hash[num] = true
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
