/**
 * Author: Deean
 * Date: 2022/11/30 23:33
 * FileName: algorithm/P1331. 数组序号转换.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	a := append([]int{}, arr...)
	sort.Ints(a)
	ranks := map[int]int{}
	for _, v := range a {
		if _, ok := ranks[v]; !ok {
			ranks[v] = len(ranks) + 1
		}
	}
	for i, v := range arr {
		arr[i] = ranks[v]
	}
	return arr
}

func main() {
	fmt.Println(arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12}))
}
