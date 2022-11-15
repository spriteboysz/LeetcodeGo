/**
 * Author: Deean
 * Date: 2022/11/15 23:52
 * FileName: algorithm/P1337. 矩阵中战斗力最弱的 K 行.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	m := len(mat)
	rows, index := make([]int, m), make([]int, m)
	for i, row := range mat {
		count := 0
		for _, v := range row {
			count += v
		}
		rows[i], index[i] = count, i
	}
	sort.Slice(index, func(i, j int) bool {
		a, b := rows[index[i]], rows[index[j]]
		return a < b || (a == b && index[i] < index[j])
	})
	return index[:k]
}

func main() {
	fmt.Println(kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}}, 3))
}
