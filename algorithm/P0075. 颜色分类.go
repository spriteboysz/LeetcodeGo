/**
 * Author: Deean
 * Date: 2022/12/10 22:26
 * FileName: algorithm/P0075. 颜色分类.go
 * Description:
 */

package main

import "fmt"

func sortColors(nums []int) {
	hash := [3]int{}
	for _, num := range nums {
		hash[num]++
	}
	color := 0
	for i := range nums {
		for hash[color] == 0 {
			color++
		}
		hash[color]--
		nums[i] = color
	}
	fmt.Println(nums)
}

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}
