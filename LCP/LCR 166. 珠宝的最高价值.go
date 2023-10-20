/**
 * Author: Deean
 * Date: 2023-10-19 22:39
 * FileName: LCR/LCR 166. 珠宝的最高价值.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func jewelleryValue(frame [][]int) int {
	Max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	m, n := len(frame), len(frame[0])
	if m == 0 || n == 0 {
		return 0
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				frame[i][j] += frame[i][j-1]
			} else if j == 0 {
				frame[i][j] += frame[i-1][j]
			} else {
				frame[i][j] += Max(frame[i-1][j], frame[i][j-1])
			}
		}
	}
	return frame[m-1][n-1]
}

func main() {
	fmt.Println(jewelleryValue(lib.Str2Array2D("[[1,3,1],[1,5,1],[4,2,1]]")))
}
