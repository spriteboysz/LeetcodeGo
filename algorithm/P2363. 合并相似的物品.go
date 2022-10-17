/**
 * Author: Deean
 * Date: 2022-10-17 23:16
 * FileName: algorithm/P2363. 合并相似的物品.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	hash := map[int]int{}
	for _, item := range items1 {
		hash[item[0]] = item[1]
	}
	for _, item := range items2 {
		hash[item[0]] += item[1]
	}
	var res [][]int
	for k, v := range hash {
		res = append(res, []int{k, v})
	}
	sort.Slice(res, func(a, b int) bool {
		return res[a][0] < res[b][0]
	})
	return res
}

func main() {
	items1 := [][]int{{1, 1}, {3, 2}, {2, 3}}
	items2 := [][]int{{2, 1}, {3, 2}, {1, 3}}
	fmt.Println(mergeSimilarItems(items1, items2))
}
