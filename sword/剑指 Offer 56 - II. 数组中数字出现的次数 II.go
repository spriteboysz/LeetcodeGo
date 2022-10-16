/**
 * Author: Deean
 * Date: 2022-10-16 16:19
 * FileName: sword/剑指 Offer 56 - II. 数组中数字出现的次数 II.go
 * Description:
 */

package main

import "fmt"

func singleNumber(nums []int) int {
	hash := make(map[int]int)
	for _, num := range nums {
		hash[num]++
	}
	for k, v := range hash {
		if v == 1 {
			return k
		}
	}
	return -1
}

func main() {
	nums := []int{9, 1, 7, 9, 7, 9, 7}
	fmt.Println(singleNumber(nums))
}
