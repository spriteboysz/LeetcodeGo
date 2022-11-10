/**
 * Author: Deean
 * Date: 2022-11-10 23:58
 * FileName: algorithm/P0999. 可以被一步捕获的棋子数.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func numRookCaptures(board [][]byte) int {
	x, y := -1, -1
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				x, y = i, j
			}
		}
	}
	rows, cols := []string{}, []string{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if i == x {
				rows = append(rows, string(board[i][j]))
			}
			if j == y {
				cols = append(cols, string(board[i][j]))
			}
		}
	}
	row, col := strings.Join(rows, ""), strings.Join(cols, "")
	row = strings.Replace(row, ".", "", -1)
	col = strings.Replace(col, ".", "", -1)
	cnt := 0
	if strings.Contains(row, "pR") {
		cnt++
	}
	if strings.Contains(row, "Rp") {
		cnt++
	}
	if strings.Contains(col, "pR") {
		cnt++
	}
	if strings.Contains(col, "Rp") {
		cnt++
	}
	return cnt
}

func main() {
	fmt.Println(numRookCaptures([][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'}}))
}
