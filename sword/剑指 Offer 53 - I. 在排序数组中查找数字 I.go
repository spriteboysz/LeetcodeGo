/**
 * Author: Deean
 * Date: 2022/11/25 22:28
 * FileName: sword/剑指 Offer 53 - I. 在排序数组中查找数字 I.go
 * Description:
 */

package main

import "fmt"

func search(nums []int, target int) int {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num]++
	}
	return hash[target]
}

func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
}
