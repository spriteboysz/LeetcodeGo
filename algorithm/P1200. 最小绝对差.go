/**
 * Author: Deean
 * Date: 2022-10-19 22:42
 * FileName: algorithm/P1200. 最小绝对差.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	minimum := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		cur := arr[i] - arr[i-1]
		if minimum > cur {
			minimum = cur
		}
	}
	var minDiff [][]int
	for i := 1; i < len(arr); i++ {
		cur := arr[i] - arr[i-1]
		if cur == minimum {
			minDiff = append(minDiff, []int{arr[i-1], arr[i]})
		}
	}
	return minDiff
}

func main() {
	arr := []int{3, 8, -10, 23, 19, -4, -14, 27}
	fmt.Println(minimumAbsDifference(arr))
}
