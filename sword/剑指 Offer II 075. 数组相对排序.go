/**
 * Author: Deean
 * Date: 2022-11-09 23:36
 * FileName: sword/剑指 Offer II 075. 数组相对排序.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	rank := map[int]int{}
	for k, v := range arr2 {
		rank[v] = k
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		rankX, hasX := rank[x]
		rankY, hasY := rank[y]
		if hasX && hasY {
			return rankX < rankY
		} else if hasX || hasY {
			return hasX
		} else {
			return x < y
		}
	})
	return arr1
}

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
		[]int{2, 1, 4, 3, 9, 6}))
}
