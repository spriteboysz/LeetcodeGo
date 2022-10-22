/**
 * Author: Deean
 * Date: 2022-10-22 15:35
 * FileName: algorithm/P0001. 两数之和.go
 * Description:
 */

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for i, num := range nums {
		if k, v := hash[target-num]; v {
			return []int{k, i}
		}
		hash[num] = i
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))
}
