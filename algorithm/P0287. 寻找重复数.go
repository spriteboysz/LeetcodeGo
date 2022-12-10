/**
 * Author: Deean
 * Date: 2022/12/9 23:14
 * FileName: algorithm/P0287. 寻找重复数.go
 * Description:
 */

package main

import "fmt"

func findDuplicate(nums []int) int {
	hash := map[int]bool{}
	for _, num := range nums {
		if hash[num] {
			return num
		}
		hash[num] = true
	}
	return -1
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}
