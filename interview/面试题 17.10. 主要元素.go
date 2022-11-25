/**
 * Author: Deean
 * Date: 2022/11/25 22:54
 * FileName: interview/面试题 17.10. 主要元素.go
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
	fmt.Println(majorityElement([]int{1, 2, 5, 9, 5, 9, 5, 5, 5}))
}
