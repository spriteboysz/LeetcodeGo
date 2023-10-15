/**
 * Author: Deean
 * Date: 2023-10-14 21:51
 * FileName: LCP/LCR 012. 寻找数组的中心下标.go
 * Description:
 */

package main

import "fmt"

func pivotIndex(nums []int) int {
	sum, acc := 0, 0
	for _, num := range nums {
		sum += num
	}
	for i, num := range nums {
		if acc*2 == sum-num {
			return i
		}
		acc += num
	}
	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
