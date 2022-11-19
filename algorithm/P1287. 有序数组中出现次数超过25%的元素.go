/**
 * Author: Deean
 * Date: 2022/11/19 19:45
 * FileName: algorithm/P1287. 有序数组中出现次数超过25%的元素.go
 * Description:
 */

package main

import "fmt"

func findSpecialInteger(arr []int) int {
	hash, n := map[int]int{}, len(arr)
	for _, num := range arr {
		hash[num]++
		if hash[num] > n/4 {
			return num
		}
	}
	return -1
}

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
}
