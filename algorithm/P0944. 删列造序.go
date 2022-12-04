/**
 * Author: Deean
 * Date: 2022/12/4 16:12
 * FileName: algorithm/P0944. 删列造序.go
 * Description:
 */

package main

import "fmt"

func minDeletionSize(strs []string) int {
	cnt := 0
	for j := range strs[0] {
		for i := 1; i < len(strs); i++ {
			if strs[i-1][j] > strs[i][j] {
				cnt++
				break
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"}))
}
