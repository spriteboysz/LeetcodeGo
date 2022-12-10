/**
 * Author: Deean
 * Date: 2022/12/10 17:07
 * FileName: algorithm/P1640. 能否连接形成数组.go
 * Description:
 */

package main

import "fmt"

func canFormArray(arr []int, pieces [][]int) bool {
	for i := 0; i < len(arr); {
		k := 0
		for k < len(pieces) && pieces[k][0] != arr[i] {
			k++
		}
		if k == len(pieces) {
			return false
		}
		for j := 0; j < len(pieces[k]) && arr[i] == pieces[k][j]; {
			i, j = i+1, j+1
		}
	}
	return true
}

func main() {
	fmt.Println(canFormArray([]int{15, 88}, [][]int{{88}, {15}}))
}
