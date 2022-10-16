/**
 * Author: Deean
 * Date: 2022-10-16 20:54
 * FileName: algorithm/P1464. 数组中两元素的最大乘积.go
 * Description:
 */

package main

import "fmt"

func maxProduct(nums []int) int {
	max1, max2 := 0, 0
	for _, num := range nums {
		if num >= max1 {
			max2 = max1
			max1 = num
		} else if num >= max2 {
			max2 = num
		}
	}
	return (max1 - 1) * (max2 - 1)
}

func main() {
	nums := []int{1, 5, 4, 5}
	fmt.Println(maxProduct(nums))
}
