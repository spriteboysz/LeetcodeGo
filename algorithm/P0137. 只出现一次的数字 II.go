/**
 * Author: Deean
 * Date: 2022-10-19 23:29
 * FileName: algorithm/P0137. 只出现一次的数字 II.go
 * Description:
 */

package main

import "fmt"

func singleNumber2(nums []int) int {
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
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber2(nums))
}
