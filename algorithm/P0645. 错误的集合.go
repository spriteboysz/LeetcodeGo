/**
 * Author: Deean
 * Date: 2022/11/25 21:47
 * FileName: algorithm/P0645. 错误的集合.go
 * Description:
 */

package main

import "fmt"

func findErrorNums(nums []int) []int {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num]++
	}
	a, b := 0, 0
	for i := 1; i <= len(nums); i++ {
		if hash[i] == 2 {
			a = i
		}
		if hash[i] == 0 {
			b = i
		}
	}
	return []int{a, b}
}

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
}
