/**
 * Author: Deean
 * Date: 2022-10-16 10:22
 * FileName: algorithm/P1913. 两个数对之间的最大乘积差.go
 * Description:
 */

package main

import "fmt"

func maxProductDifference(nums []int) int {
	max1, max2, min1, min2 := 0, 0, 10001, 10001
	for _, num := range nums {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
		if num < min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}
	}
	return max1*max2 - min1*min2
}

func main() {
	nums := []int{4, 2, 5, 9, 7, 4, 8}
	fmt.Println(maxProductDifference(nums))
}
