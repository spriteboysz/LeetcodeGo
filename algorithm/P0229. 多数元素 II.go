/**
 * Author: Deean
 * Date: 2022/12/14 22:38
 * FileName: algorithm/P0229. 多数元素 II.go
 * Description:
 */

package main

import "fmt"

func majorityElementII(nums []int) []int {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num]++
	}
	majority := []int{}
	for k, v := range hash {
		if v > len(nums)/3 {
			majority = append(majority, k)
		}
	}
	return majority
}

func main() {
	fmt.Println(majorityElementII([]int{1, 2}))
}
