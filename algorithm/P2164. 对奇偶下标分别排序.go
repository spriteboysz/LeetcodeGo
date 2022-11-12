/**
 * Author: Deean
 * Date: 2022/11/12 20:16
 * FileName: algorithm/P2164. 对奇偶下标分别排序.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func sortEvenOdd(nums []int) []int {
	even, odd := []int{}, []int{}
	for i, num := range nums {
		if i%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	sort.Ints(even)
	sort.Ints(odd)
	sorted := []int{}
	for i := 0; i < len(even); i++ {
		sorted = append(sorted, even[i])
		if i < len(odd) {
			sorted = append(sorted, odd[len(odd)-i-1])
		}
	}
	return sorted
}

func main() {
	fmt.Println(sortEvenOdd([]int{4, 1, 2, 3}))
}
