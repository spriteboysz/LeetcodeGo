/**
 * Author: Deean
 * Date: 2022-10-20 23:49
 * FileName: algorithm/P0169. 多数元素.go
 * Description:
 */

package main

import "fmt"

func majorityElement(nums []int) int {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num]++
		if hash[num] > len(nums)/2 {
			return num
		}
	}
	return -1
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
