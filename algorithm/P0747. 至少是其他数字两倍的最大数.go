/**
 * Author: Deean
 * Date: 2022/11/23 23:21
 * FileName: algorithm/P0747. 至少是其他数字两倍的最大数.go
 * Description:
 */

package main

import "fmt"

func dominantIndex(nums []int) int {
	max1, max2, index := -1, -1, 0
	for i, num := range nums {
		if num > max1 {
			max1, max2, index = num, max1, i
		} else if num > max2 {
			max2 = num
		}
	}
	if max1 >= max2*2 {
		return index
	}
	return -1
}

func main() {
	fmt.Println(dominantIndex([]int{1, 2, 3, 4}))
}
