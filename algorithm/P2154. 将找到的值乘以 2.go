/**
 * Author: Deean
 * Date: 2022/11/14 23:18
 * FileName: algorithm/P2154. 将找到的值乘以 2.go
 * Description:
 */

package main

import "fmt"

func findFinalValue(nums []int, original int) int {
	hash := map[int]bool{}
	for _, num := range nums {
		hash[num] = true
	}
	for hash[original] {
		original *= 2
	}
	return original
}

func main() {
	fmt.Println(findFinalValue([]int{5, 3, 6, 1, 12}, 3))
}
