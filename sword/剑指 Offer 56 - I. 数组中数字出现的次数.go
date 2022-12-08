/**
 * Author: Deean
 * Date: 2022/12/7 23:09
 * FileName: sword/剑指 Offer 56 - I. 数组中数字出现的次数.go
 * Description:
 */

package main

import "fmt"

func singleNumbers(nums []int) []int {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num]++
	}
	single := []int{}
	for k, v := range hash {
		if v == 1 {
			single = append(single, k)
		}
	}
	return single
}

func main() {
	fmt.Println(singleNumbers([]int{4, 1, 4, 6}))
}
