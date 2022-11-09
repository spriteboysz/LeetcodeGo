/**
 * Author: Deean
 * Date: 2022-11-09 22:55
 * FileName: algorithm/P0905. 按奇偶排序数组.go
 * Description:
 */

package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	odd, even := []int{}, []int{}
	for _, num := range nums {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	return append(even, odd...)
}

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity([]int{3, 1}))
	fmt.Println(sortArrayByParity([]int{2, 4}))
}
