/**
 * Author: Deean
 * Date: 2022-10-20 23:34
 * FileName: algorithm/P0961. 在长度 2N 的数组中找出重复 N 次的元素.go
 * Description:
 */

package main

import "fmt"

func repeatedNTimes(nums []int) int {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num]++
		if hash[num] > 1 {
			return num
		}
	}
	return -1
}

func main() {
	nums := []int{5, 1, 5, 2, 5, 3, 5, 4}
	fmt.Println(repeatedNTimes(nums))
}
