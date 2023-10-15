/**
 * Author: Deean
 * Date: 2023-10-14 20:54
 * FileName: LCP/LCR 004. 只出现一次的数字 II.go
 * Description:
 */

package main

import "fmt"

func singleNumber(nums []int) int {
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
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 100}))
}
