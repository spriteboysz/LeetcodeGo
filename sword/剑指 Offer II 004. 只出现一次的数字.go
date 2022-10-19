/**
 * Author: Deean
 * Date: 2022-10-20 00:01
 * FileName: algorithm/剑指 Offer II 004. 只出现一次的数字.go
 * Description:
 */

package main

import "fmt"

func singleNumber3(nums []int) int {
	hash := map[int]int{}
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
	nums := []int{0, 1, 0, 1, 0, 1, 100}
	fmt.Println(singleNumber3(nums))
}
