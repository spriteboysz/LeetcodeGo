/**
 * Author: Deean
 * Date: 2022/11/12 20:31
 * FileName: algorithm/P1582. 二进制矩阵中的特殊位置.go
 * Description:
 */

package main

import "fmt"

func numSpecial(mat [][]int) int {
	rows, cols := make([]int, len(mat)), make([]int, len(mat[0]))
	for i, row := range mat {
		for j, num := range row {
			rows[i] += num
			cols[j] += num
		}
	}
	cnt := 0
	for i, row := range mat {
		for j, num := range row {
			if num == 1 && rows[i] == 1 && cols[j] == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(numSpecial([][]int{
		{0, 0, 0, 1},
		{1, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0}}))
}
